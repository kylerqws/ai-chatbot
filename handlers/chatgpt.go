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

// GetChatGPTResponse sends a request to OpenAI and retrieves the response
func GetChatGPTResponse(prompt, apiKey string) (string, error) {
	// Create the request body for Chat Completions
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo", // Specify the model to use
		"messages": []map[string]string{
			{"role": "user", "content": prompt}, // Role: user, Content: the user's prompt
		},
		"max_tokens": 150,   // Maximum length of the response
		"temperature": 0.7,  // Controls randomness/creativity of the response
	})
	if err != nil {
		return "", fmt.Errorf("error creating request body: %v", err)
	}

	// Create an HTTP POST request
	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey) // Add API key for authentication
	req.Header.Set("Content-Type", "application/json") // Specify that the body is JSON

	// HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Timeout of 10 seconds
	}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	// Handle errors from OpenAI API
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error from OpenAI (status: %d): %s", resp.StatusCode, string(body))
	}

	// Decode the JSON response
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	// Extract the text response
	choices, ok := response["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("empty or invalid choices in response")
	}

	// Trim whitespace and return the result
	result := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return strings.TrimSpace(result), nil
}
