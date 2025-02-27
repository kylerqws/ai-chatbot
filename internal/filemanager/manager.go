package filemanager

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// Manager управляет файлами в OpenAI API.
type Manager struct {
	client *Client
}

// NewManager создаёт новый экземпляр Manager.
func NewManager(apiKey string) *Manager {
	client := NewClient(apiKey)
	return &Manager{client: client}
}

// // UploadFile загружает файл в OpenAI и возвращает file ID.
// func (m *Manager) UploadFile(filePath, purpose string) (string, error) {
// 	body := map[string]string{
// 		"purpose": purpose,
// 	}

// 	response, err := m.client.SendRequest("POST", uploadURL, body)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to upload file: %v", err)
// 	}

// 	fileID, exists := response["id"].(string)
// 	if !exists {
// 		return "", fmt.Errorf("file ID not found in response")
// 	}

// 	return fileID, nil
// }

func (m *Manager) UploadFile(filePath, purpose string) (string, error) {
	// Конвертируем JSON в JSONL (если нужно)
	fileReader, newFileName, err := ConvertJSONtoJSONL(filePath)
	if err != nil {
		return "", fmt.Errorf("conversion failed: %v", err)
	}

	// Создаём multipart body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Добавляем поле "purpose"
	_ = writer.WriteField("purpose", purpose)

	// Добавляем файл (с именем после конвертации)
	part, err := writer.CreateFormFile("file", newFileName)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}
	_, err = io.Copy(part, fileReader)
	if err != nil {
		return "", fmt.Errorf("failed to copy file data: %v", err)
	}

	// Завершаем запись multipart
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// Создаём HTTP-запрос
	req, err := http.NewRequest("POST", uploadURL, &requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Authorization", "Bearer "+m.client.apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Отправляем запрос
	resp, err := m.client.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %s", bodyResp)
	}

	// Возвращаем ID загруженного файла
	return string(bodyResp), nil
}

// GetFileInfo получает информацию о файле.
func (m *Manager) GetFileInfo(fileID string) (map[string]interface{}, error) {
	url := fmt.Sprintf(fileURL, fileID)
	return m.client.SendRequest("GET", url, nil)
}

// DeleteFile удаляет файл по его ID.
func (m *Manager) DeleteFile(fileID string) error {
	url := fmt.Sprintf(fileURL, fileID)

	_, err := m.client.SendRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}

// DeleteAllFiles удаляет все загруженные файлы.
func (m *Manager) DeleteAllFiles() error {
	files, err := m.ListFiles()
	if err != nil {
		return fmt.Errorf("failed to fetch file list: %v", err)
	}

	for _, file := range files {
		fileID, _ := file["id"].(string)

		err := m.DeleteFile(fileID)
		if err != nil {
			fmt.Printf("Failed to delete file %s: %v\n", fileID, err)
		} else {
			fmt.Printf("File %s deleted successfully\n", fileID)
		}
	}
	return nil
}

// ListFiles получает список всех загруженных файлов.
func (m *Manager) ListFiles() ([]map[string]interface{}, error) {
	response, err := m.client.SendRequest("GET", uploadURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file list: %v", err)
	}

	files, exists := response["data"].([]interface{})
	if !exists {
		return nil, fmt.Errorf("unexpected response format")
	}

	var fileList []map[string]interface{}
	for _, f := range files {
		if fileData, ok := f.(map[string]interface{}); ok {
			fileList = append(fileList, fileData)
		}
	}

	return fileList, nil
}
