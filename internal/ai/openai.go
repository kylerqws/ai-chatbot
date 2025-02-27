package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

// OpenAIClient - клиент OpenAI.
type OpenAIClient struct {
	apiKey string
	client *http.Client
}

// NewOpenAIClient создаёт OpenAI клиента.
func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		apiKey: apiKey,
		client: &http.Client{Timeout: 60 * time.Second},
	}
}

// GetResponse отправляет запрос в OpenAI и получает ответ.
func (o *OpenAIClient) GetResponse(prompt string) (string, error) {
	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	})

	req, _ := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+o.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := o.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("OpenAI request failed: %v", err)
	}
	defer resp.Body.Close()

	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	json.NewDecoder(resp.Body).Decode(&response)

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("empty response from OpenAI")
	}

	return response.Choices[0].Message.Content, nil
}
