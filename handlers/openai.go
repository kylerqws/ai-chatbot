package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

func GetOpenAIResponse(prompt, apiKey string) (string, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
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

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error executing request: %v", err)
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
	return strings.TrimSpace(result), nil
}
