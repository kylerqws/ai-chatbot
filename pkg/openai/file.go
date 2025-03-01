package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/kylerqws/chatgpt-bot/pkg/converter"
)

const (
	uploadURL = "https://api.openai.com/v1/files"
	fileURL   = "https://api.openai.com/v1/files/%s"
)

// FileClient выполняет запросы к API OpenAI для работы с файлами.
type FileClient struct {
	apiKey string
	client *http.Client
}

// NewFileClient создаёт новый экземпляр File API Client.
func NewFileClient(apiKey string) *FileClient {
	return &FileClient{
		apiKey: apiKey,
		client: &http.Client{Timeout: 60 * time.Second},
	}
}

// UploadFile загружает файл в OpenAI и возвращает file ID.
func (fc *FileClient) UploadFile(filePath, purpose string) (string, error) {
	// Конвертируем JSON в JSONL, если нужно
	fileReader, _, err := converter.ConvertJSONtoJSONL(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to convert file: %v", err)
	}

	// Создаём multipart-запрос
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}
	if _, err := io.Copy(part, fileReader); err != nil {
		return "", fmt.Errorf("failed to copy file content: %v", err)
	}
	writer.WriteField("purpose", purpose)
	writer.Close()

	// Отправляем запрос
	req, err := http.NewRequest("POST", uploadURL, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+fc.apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := fc.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("file upload failed: %v", err)
	}
	defer resp.Body.Close()

	// Декодируем ответ
	var response struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return response.ID, nil
}

// GetFileInfo получает информацию о файле по его ID.
func (fc *FileClient) GetFileInfo(fileID string) (map[string]interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf(fileURL, fileID), nil)
	req.Header.Set("Authorization", "Bearer "+fc.apiKey)

	resp, err := fc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}
	defer resp.Body.Close()

	var fileInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&fileInfo); err != nil {
		return nil, fmt.Errorf("failed to decode file info response: %v", err)
	}

	return fileInfo, nil
}

// DeleteFile удаляет файл по его ID.
func (fc *FileClient) DeleteFile(fileID string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf(fileURL, fileID), nil)
	req.Header.Set("Authorization", "Bearer "+fc.apiKey)

	resp, err := fc.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}
	defer resp.Body.Close()

	return nil
}

// DeleteAllFiles удаляет все файлы из OpenAI.
func (fc *FileClient) DeleteAllFiles() error {
	files, err := fc.ListFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		fileID, ok := file["id"].(string)
		if !ok {
			continue
		}
		_ = fc.DeleteFile(fileID)
	}

	return nil
}

// ListFiles получает список всех загруженных файлов.
func (fc *FileClient) ListFiles() ([]map[string]interface{}, error) {
	req, _ := http.NewRequest("GET", uploadURL, nil)
	req.Header.Set("Authorization", "Bearer "+fc.apiKey)

	resp, err := fc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %v", err)
	}
	defer resp.Body.Close()

	var response struct {
		Data []map[string]interface{} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode list files response: %v", err)
	}

	return response.Data, nil
}
