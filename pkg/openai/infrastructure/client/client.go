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
	return &Client{
		config:     cfg,
		httpClient: &http.Client{Timeout: cfg.GetTimeout()},
	}
}

func (c *Client) RequestMultipart(ctx context.Context, path string, body map[string]string) ([]byte, error) {
	filePath := body["file"]
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("client: failed to open file %q: %w", filePath, err)
	}
	defer func(file *os.File) {
		if clerr := file.Close(); clerr != nil && err == nil {
			err = fmt.Errorf("client: failed to close file %q: %w", filePath, clerr)
		}
	}(file)

	var fileReader io.Reader = file
	if strings.HasSuffix(strings.ToLower(filePath), ".json") {
		prp := body["purpose"]
		if prp == purpose.FineTune.Code || prp == purpose.FineTuneResults.Code {
			fileReader, err = jsonl.ConvertToReader(filePath)
			if err != nil {
				return nil, fmt.Errorf("client: failed to convert json to jsonl: %w", err)
			}
		}
	}

	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)

	part, err := w.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("client: failed to create multipart file part: %w", err)
	}
	if _, err = io.Copy(part, fileReader); err != nil {
		return nil, fmt.Errorf("client: failed to copy file content: %w", err)
	}

	for key, val := range body {
		if key != "file" {
			if err := w.WriteField(key, val); err != nil {
				return nil, fmt.Errorf("client: failed to write multipart field %q: %w", key, err)
			}
		}
	}

	if err = w.Close(); err != nil {
		return nil, fmt.Errorf("client: failed to close multipart writer: %w", err)
	}

	req, err := c.buildRequest("POST", path, b)
	if err != nil {
		return nil, fmt.Errorf("client: failed to build request: %w", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return c.doRequest(ctx, req)
}

func (c *Client) RequestJSON(ctx context.Context, method, path string, body any) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, fmt.Errorf("client: failed to encode json body: %w", err)
	}

	req, err := c.buildRequest(method, path, buf)
	if err != nil {
		return nil, fmt.Errorf("client: failed to build request: %w", err)
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
		return nil, fmt.Errorf("client: failed to build request: %w", err)
	}
	return c.doRequest(ctx, req)
}

func (c *Client) buildRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := strings.TrimRight(c.config.GetBaseURL(), "/") + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("client: failed to create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.config.GetAPIKey())

	return req, nil
}

func (c *Client) doRequest(ctx context.Context, req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("client: HTTP request failed: %w", err)
	}
	defer func(body io.ReadCloser) {
		if clerr := body.Close(); clerr != nil && err == nil {
			err = fmt.Errorf("client: failed to close response body: %w", clerr)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("client: failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		msg := c.extractAPIError(body)
		return nil, fmt.Errorf("client: unexpected status %s (%s)", resp.Status, msg)
	}

	return body, nil
}

func (c *Client) extractAPIError(body []byte) string {
	var data struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(body, &data); err == nil && data.Error.Message != "" {
		return data.Error.Message
	}

	return "unknown OpenAI API error"
}
