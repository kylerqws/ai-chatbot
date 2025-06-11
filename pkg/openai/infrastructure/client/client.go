package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/kylerqws/chatbot/pkg/openai/enumset/purpose"
	"github.com/kylerqws/chatbot/pkg/openai/utils/converter/jsonl"

	ctrcl "github.com/kylerqws/chatbot/pkg/openai/contract/client"
	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
)

// client implements the Client interface for communicating with the OpenAI API.
type client struct {
	config     ctrcfg.Config
	httpClient *http.Client
}

// New creates a new OpenAI API client using the provided configuration.
func New(cfg ctrcfg.Config) ctrcl.Client {
	hc := &http.Client{Timeout: cfg.GetTimeout()}
	return &client{config: cfg, httpClient: hc}
}

// RequestMultipart sends a multipart/form-data POST request to the specified path.
// The body map must include a "file" entry. If the file is JSON and the purpose is for fine-tuning,
// it will be automatically converted to JSONL format.
func (c *client) RequestMultipart(ctx context.Context, path string, body map[string]string) (res []byte, err error) {
	filePath := body["file"]
	if filePath == "" {
		return nil, fmt.Errorf("missing required field 'file' in request body")
	}

	reader, err := c.prepareFileReader(filePath, body["purpose"])
	if err != nil {
		return nil, fmt.Errorf("prepare file reader: %w", err)
	}
	defer func() {
		if closer, ok := reader.(io.Closer); ok {
			if cerr := closer.Close(); cerr != nil && err == nil {
				err = fmt.Errorf("close file reader: %w", cerr)
			}
		}
	}()

	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	go c.writeMultipartBody(writer, pw, reader, filePath, body)

	req, err := c.buildRequest(ctx, http.MethodPost, path, pr)
	if err != nil {
		return nil, fmt.Errorf("build multipart request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return c.doRequest(req)
}

// RequestJSON sends an HTTP request with a JSON-encoded body.
func (c *client) RequestJSON(ctx context.Context, method, path string, body any) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, fmt.Errorf("encode JSON body: %w", err)
	}

	req, err := c.buildRequest(ctx, method, path, buf)
	if err != nil {
		return nil, fmt.Errorf("build JSON request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return c.doRequest(req)
}

// RequestRaw sends a generic HTTP request with the provided method, path, and optional body.
// Suitable for GET, DELETE, etc.
func (c *client) RequestRaw(ctx context.Context, method, path string, body io.Reader) ([]byte, error) {
	req, err := c.buildRequest(ctx, method, path, body)
	if err != nil {
		return nil, fmt.Errorf("build raw request: %w", err)
	}
	return c.doRequest(req)
}

// prepareFileReader returns a file reader or a JSONL-converted stream based on file and purpose.
func (c *client) prepareFileReader(filePath, purposeCode string) (io.ReadCloser, error) {
	if jsonl.HasJSONSuffix(filePath) && purposeCode == purpose.FineTune.Code {
		r, err := jsonl.ConvertToReader(filePath)
		if err != nil {
			return nil, fmt.Errorf("convert to JSONL: %w", err)
		}
		return r, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file '%s': %w", filePath, err)
	}
	return file, nil
}

// buildRequest constructs a new HTTP request with appropriate headers and context.
func (c *client) buildRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	url := strings.TrimRight(c.config.GetBaseURL(), "/") + "/" + strings.TrimLeft(path, "/")
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.config.GetAPIKey())
	return req, nil
}

// doRequest performs the HTTP request and handles the response.
func (c *client) doRequest(req *http.Request) (body []byte, err error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close response body: %w", cerr)
		}
	}()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("OpenAI API error: %s", extractAPIError(body))
	}

	return body, nil
}

// writeMultipartBody writes the file and additional fields to the multipart writer.
func (c *client) writeMultipartBody(writer *multipart.Writer, pw *io.PipeWriter, reader io.Reader, filePath string, fields map[string]string) {
	var err error
	defer func() {
		if cerr := writer.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close multipart writer: %w", cerr)
		}
		if cerr := pw.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close pipe writer: %w", cerr)
		}
	}()

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		_ = pw.CloseWithError(fmt.Errorf("create form file: %w", err))
		return
	}
	if _, err := io.Copy(part, reader); err != nil {
		_ = pw.CloseWithError(fmt.Errorf("copy file content: %w", err))
		return
	}
	for k, v := range fields {
		if k == "file" {
			continue
		}
		if err := writer.WriteField(k, v); err != nil {
			_ = pw.CloseWithError(fmt.Errorf("write field '%s': %w", k, err))
			return
		}
	}
}

// extractAPIError parses a structured OpenAI error response and formats it.
func extractAPIError(body []byte) string {
	var data struct {
		Error struct {
			Message string  `json:"message"`
			Type    *string `json:"type,omitempty"`
			Param   *string `json:"param,omitempty"`
			Code    *string `json:"code,omitempty"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &data); err == nil && data.Error.Message != "" {
		msg := data.Error.Message
		var parts []string
		if t := data.Error.Type; t != nil && *t != "" {
			parts = append(parts, fmt.Sprintf("type '%s'", *t))
		}
		if p := data.Error.Param; p != nil && *p != "" {
			parts = append(parts, fmt.Sprintf("param '%s'", *p))
		}
		if c := data.Error.Code; c != nil && *c != "" {
			parts = append(parts, fmt.Sprintf("code '%s'", *c))
		}
		if len(parts) > 0 {
			msg += " (" + strings.Join(parts, ", ") + ")"
		}
		return msg
	}
	return "unknown OpenAI API error"
}
