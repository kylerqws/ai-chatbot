package filemanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// ConvertJSONtoJSONL читает JSON-файл и преобразует его в JSONL-формат (налету).
func ConvertJSONtoJSONL(filePath string) (io.Reader, string, error) {
	// Проверяем, что файл имеет расширение .json
	if !strings.HasSuffix(filePath, ".json") {
		// Если это не JSON, просто возвращаем сам файл как поток данных
		file, err := os.Open(filePath)
		if err != nil {
			return nil, "", fmt.Errorf("failed to open file: %v", err)
		}
		return file, filePath, nil
	}

	// Читаем JSON-файл
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read JSON file: %v", err)
	}

	// Декодируем JSON
	var records []map[string]interface{}
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, "", fmt.Errorf("invalid JSON format: %v", err)
	}

	// Создаём буфер для хранения JSONL-данных
	var buffer bytes.Buffer

	// Записываем каждую строку в JSONL-формате в буфер
	for _, record := range records {
		line, err := json.Marshal(record)
		if err != nil {
			return nil, "", fmt.Errorf("failed to encode JSONL line: %v", err)
		}
		buffer.Write(line)
		buffer.WriteByte('\n')
	}

	// Возвращаем буфер в виде io.Reader
	return &buffer, strings.TrimSuffix(filePath, ".json") + ".jsonl", nil
}
