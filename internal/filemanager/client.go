package filemanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// API endpoints
const (
	uploadURL = "https://api.openai.com/v1/files"
	fileURL   = "https://api.openai.com/v1/files/%s"
)

// Client выполняет запросы к API OpenAI для работы с файлами.
type Client struct {
	apiKey string
	client *http.Client
}

// NewClient создаёт новый экземпляр File API Client.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		client: &http.Client{Timeout: 60 * time.Second},
	}
}

// SendRequest выполняет HTTP-запрос к API.
func (c *Client) SendRequest(method, url string, body interface{}) (map[string]interface{}, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to encode request body: %v", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", bodyResp)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(bodyResp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return response, nil
}
