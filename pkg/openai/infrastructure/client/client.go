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

	"github.com/kylerqws/chatbot/pkg/openai/domain/purpose"
	"github.com/kylerqws/chatbot/pkg/openai/utils/converter/jsonl"

	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
)

type Client struct {
	config     ctrcfg.Config
	httpClient *http.Client
}

func New(cfg ctrcfg.Config) *Client {
	return &Client{config: cfg, httpClient: &http.Client{Timeout: cfg.GetTimeout()}}
}

func (c *Client) RequestMultipart(ctx context.Context, path string, body map[string]string) (resp []byte, err error) {
	filePath := body["file"]
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("[client.RequestMultipart] failed to open file %v: %w", filePath, err)
	}
	defer func(file *os.File) {
		if clerr := file.Close(); clerr != nil {
			if err == nil {
				err = fmt.Errorf("[client.RequestMultipart] failed to close file %v: %w", filePath, clerr)
			}
		}
	}(file)

	reader := io.Reader(file)
	if strings.HasSuffix(strings.ToLower(filePath), ".json") {
		prp := body["purpose"]
		if prp == purpose.FineTune.Code {
			if reader, err = jsonl.ConvertToReader(filePath); err != nil {
				return nil, fmt.Errorf("[client.RequestMultipart] failed to convert json to jsonl: %w", err)
			}
		}
	}

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	err = c.writeMultipart(writer, reader, filepath.Base(filePath), body)
	if err != nil {
		return nil, fmt.Errorf("[client.RequestMultipart] failed to write multipart: %w", err)
	}

	req, err := c.buildRequest("POST", path, buf)
	if err != nil {
		return nil, fmt.Errorf("[client.RequestMultipart] failed to build request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err = c.doRequest(ctx, req)
	return resp, err
}

func (c *Client) RequestJSON(ctx context.Context, method, path string, body any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, fmt.Errorf("[client.RequestJSON] failed to encode json body: %w", err)
	}

	req, err := c.buildRequest(method, path, buf)
	if err != nil {
		return nil, fmt.Errorf("[client.RequestJSON] failed to build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return c.doRequest(ctx, req)
}

func (c *Client) Request(ctx context.Context, method, path string) ([]byte, error) {
	return c.RequestReader(ctx, method, path, nil)
}

func (c *Client) RequestReader(ctx context.Context, method, path string, body io.Reader) ([]byte, error) {
	req, err := c.buildRequest(method, path, body)
	if err != nil {
		return nil, fmt.Errorf("[client.RequestReader] failed to build request: %w", err)
	}

	return c.doRequest(ctx, req)
}

func (c *Client) buildRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := strings.TrimRight(c.config.GetBaseURL(), "/") + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("[client.buildRequest] failed to create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.config.GetAPIKey())

	return req, nil
}

func (_ *Client) writeMultipart(w *multipart.Writer, file io.Reader, filename string, fields map[string]string) error {
	part, err := w.CreateFormFile("file", filename)
	if err != nil {
		return fmt.Errorf("[client.writeMultipart] failed to create multipart file part: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("[client.writeMultipart] failed to copy file content: %w", err)
	}

	for k, v := range fields {
		if k != "file" {
			err = w.WriteField(k, v)
			if err != nil {
				return fmt.Errorf("[client.writeMultipart] failed to write field '%v': %w", k, err)
			}
		}
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("[client.writeMultipart] failed to close multipart writer: %w", err)
	}

	return nil
}

func (c *Client) doRequest(ctx context.Context, req *http.Request) (body []byte, err error) {
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("[client.doRequest] HTTP request failed: %w", err)
	}
	defer func(body io.ReadCloser) {
		if clerr := body.Close(); clerr != nil {
			if err == nil {
				err = fmt.Errorf("[client.doRequest] failed to close response body: %w", clerr)
			}
		}
	}(resp.Body)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[client.doRequest] failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		msg := c.extractAPIError(body)
		return nil, fmt.Errorf("[client.doRequest] unexpected status '%v' (%s)", resp.Status, msg)
	}

	return body, err
}

func (_ *Client) extractAPIError(body []byte) string {
	var data struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}

	err := json.Unmarshal(body, &data)
	if err == nil && data.Error.Message != "" {
		return data.Error.Message
	}

	return "unknown OpenAI API error"
}
