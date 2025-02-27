package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kylerqws/go-chatgpt-vk/models"
)

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

var httpClient = &http.Client{
	Timeout: 60 * time.Second,
}

func GetOpenAIResponse(prompt string, apiKey string) (string, error) {
	modelData, err := models.ModelConfig.Load()
	if err != nil {
		return "", fmt.Errorf("Failed to load model config: %v", err)
	}
	modelName := modelData.Name

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": modelName,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"max_tokens":  150,
		"temperature": 0.7,
	})
	if err != nil {
		return "", fmt.Errorf("Error creating request body: %v", err)
	}

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Error from OpenAI (status: %d): %s", resp.StatusCode, string(body))
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("Error decoding response: %v", err)
	}

	choices, ok := response["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("Empty or invalid choices in response")
	}

	result := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return result, nil
}
