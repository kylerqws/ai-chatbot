package training

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
	trainURL  = "https://api.openai.com/v1/fine_tuning/jobs"
	jobURL    = "https://api.openai.com/v1/fine_tuning/jobs/%s"
	cancelURL = "https://api.openai.com/v1/fine_tuning/jobs/%s/cancel"
)

// OpenAIClient - клиент для взаимодействия с OpenAI API.
type OpenAIClient struct {
	apiKey string
	client *http.Client
}

// NewOpenAIClient создает новый экземпляр OpenAI API клиента.
func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		apiKey: apiKey,
		client: &http.Client{Timeout: 60 * time.Second},
	}
}

// SendRequest отправляет HTTP-запрос к OpenAI API и возвращает ответ.
func (c *OpenAIClient) SendRequest(method, url string, body interface{}) (map[string]interface{}, error) {
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
