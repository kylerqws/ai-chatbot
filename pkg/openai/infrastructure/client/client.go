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
		return nil, err
	}
	defer func(file *os.File) {
		if clerr := file.Close(); clerr != nil && err == nil {
			err = clerr
		}
	}(file)

	var fileReader io.Reader = file
	if strings.HasSuffix(strings.ToLower(filePath), ".json") {
		prp := body["purpose"]
		if prp == purpose.FineTune.Code || prp == purpose.FineTuneResults.Code {
			fileReader, err = jsonl.ConvertToReader(filePath)
			if err != nil {
				return nil, err
			}
		}
	}

	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)

	part, err := w.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(part, fileReader); err != nil {
		return nil, err
	}

	for key, val := range body {
		if key != "file" {
			if err := w.WriteField(key, val); err != nil {
				return nil, err
			}
		}
	}
	if err = w.Close(); err != nil {
		return nil, err
	}

	req, err := c.buildRequest("POST", path, b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return c.doRequest(ctx, req)
}

func (c *Client) RequestJSON(cfg context.Context, method, path string, body any) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, err
	}

	req, err := c.buildRequest(method, path, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return c.doRequest(cfg, req)
}

func (c *Client) Request(cfg context.Context, method, path string) ([]byte, error) {
	return c.RequestReader(cfg, method, path, nil)
}

func (c *Client) RequestReader(cfg context.Context, method, path string, body io.Reader) ([]byte, error) {
	req, err := c.buildRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	return c.doRequest(cfg, req)
}

func (c *Client) buildRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := strings.TrimRight(c.config.GetBaseURL(), "/") + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.config.GetAPIKey())

	return req, nil
}

func (c *Client) doRequest(ctx context.Context, req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer func(body io.ReadCloser) {
		if clerr := body.Close(); clerr != nil && err == nil {
			err = clerr
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		msg := c.extractAPIError(body)
		return nil, fmt.Errorf("%s (%s)", resp.Status, msg)
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

	return "OpenAI API error"
}
