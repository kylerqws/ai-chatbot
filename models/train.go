package models

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	uploadURL = "https://api.openai.com/v1/files"
	deleteURL = "https://api.openai.com/v1/files/%s"
	trainURL  = "https://api.openai.com/v1/fine_tuning/jobs"
	purpose   = "fine-tune"
)

func convertJSONToJSONL(inputPath string) (*bytes.Buffer, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open JSON file: %v", err)
	}
	defer file.Close()

	var data []map[string]interface{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse JSON: %v", err)
	}

	jsonlBuffer := &bytes.Buffer{}
	writer := bufio.NewWriter(jsonlBuffer)

	for _, entry := range data {
		jsonData, err := json.Marshal(entry)
		if err != nil {
			return nil, fmt.Errorf("Failed to encode JSON line: %v", err)
		}
		_, err = writer.WriteString(string(jsonData) + "\n")
		if err != nil {
			return nil, fmt.Errorf("Failed to write JSONL data: %v", err)
		}
	}
	writer.Flush()
	return jsonlBuffer, nil
}

func deleteFileFromOpenAI(apiKey, fileID string) error {
	if fileID == "" {
		return nil
	}

	url := fmt.Sprintf(deleteURL, fileID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("Failed to create delete request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to send delete request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Failed to delete file: %s", string(body))
	}

	log.Printf("Previous file %s deleted successfully from OpenAI.", fileID)
	return nil
}

func uploadTrainingFile(apiKey string, previousFileID string) (string, error) {
	if err := deleteFileFromOpenAI(apiKey, previousFileID); err != nil {
		log.Printf("Warning: Could not delete previous file: %v", err)
	}

	inputFile := "config/prompts.json"
	jsonlBuffer, err := convertJSONToJSONL(inputFile)
	if err != nil {
		return "", fmt.Errorf("Error converting JSON to JSONL: %v", err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "prompts.jsonl")
	if err != nil {
		return "", fmt.Errorf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, jsonlBuffer)
	if err != nil {
		return "", fmt.Errorf("Failed to write JSONL data to form: %v", err)
	}

	_ = writer.WriteField("purpose", "fine-tune")
	writer.Close()

	req, err := http.NewRequest("POST", uploadURL, body)
	if err != nil {
		return "", fmt.Errorf("Failed to create upload request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to send upload request: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("File Upload Failed: %s", string(bodyBytes))
	}

	var response map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return "", fmt.Errorf("Failed to parse upload response: %v", err)
	}

	fileID, exists := response["id"].(string)
	if !exists {
		return "", fmt.Errorf("File ID not found in response")
	}
	return fileID, nil
}

func trainModel(apiKey, fileID string) (string, error) {
	if fileID == "" {
		return "", fmt.Errorf("File ID is empty. Training cannot proceed.")
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"training_file": fileID,
		"model":         "gpt-3.5-turbo",
		"suffix":        "custom_socionics_model",
	})
	if err != nil {
		return "", fmt.Errorf("Failed to create training request body: %v", err)
	}

	req, err := http.NewRequest("POST", trainURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("Failed to create training request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to send training request: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Training Failed: %s", string(bodyBytes))
	}

	var response map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return "", fmt.Errorf("Failed to parse response: %v\nResponse body: %s", err, string(bodyBytes))
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

	var previousFileID string
	if modelData, err := ModelConfig.Load(); err == nil {
		previousFileID = modelData.FileID
	}

	fileID, err := uploadTrainingFile(apiKey, previousFileID)
	if err != nil {
		log.Fatalf("File Upload Failed: %v\n", err)
	}

	log.Printf("File uploaded successfully. File ID: %s", fileID)

	modelName, err := trainModel(apiKey, fileID)
	if err != nil {
		log.Fatalf("Fine-Tuning Failed: %v\n", err)
	}

	newModelData := ModelData{Name: modelName, FileID: fileID}
	if err := ModelConfig.Save(&newModelData); err != nil {
		log.Fatalf("Failed To Save Model Name: %v\n", err)
	}

	fmt.Println("Model Trained Successfully:", modelName)
}
