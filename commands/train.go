package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kylerqws/go-chatgpt-vk/models"
)

const trainURL = "https://api.openai.com/v1/fine-tunes"

type ModelData struct {
	Name string `json:"name"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func loadPrompts() ([]byte, error) {
	var prompts []map[string]interface{}
	err := models.PromptsConfig.Load(&prompts)
	if err != nil {
		return nil, fmt.Errorf("Failed to load prompts: %v", err)
	}

	data, err := json.Marshal(prompts)
	if err != nil {
		return nil, fmt.Errorf("Failed to encode prompts: %v", err)
	}

	return data, nil
}

func trainModel(apiKey string) (string, error) {
	prompts, err := loadPrompts()
	if err != nil {
		return "", err
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"training_file": string(prompts),
		"model":         "gpt-4",
	})
	if err != nil {
		return "", fmt.Errorf("Failed to create request body: %v", err)
	}

	req, err := http.NewRequest("POST", trainURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("Failed to create HTTP request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("Failed to parse response: %v", err)
	}

	if id, exists := response["id"].(string); exists {
		return id, nil
	}
	return "", fmt.Errorf("Error: Model ID not found in response")
}

func Train() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("API Key Not Found. Ensure OPENAI_API_KEY Is Set In .env")
	}

	modelName, err := trainModel(apiKey)
	if err != nil {
		log.Fatalf("Fine-Tuning Failed: %v\n", err)
	}

	modelData := ModelData{Name: modelName}
	if err := models.ModelConfig.Save(&modelData); err != nil {
		log.Fatalf("Failed To Save Model Name: %v\n", err)
	}

	fmt.Println("Model Trained Successfully:", modelName)
}
