package converter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// ConvertJSONtoJSONL читает JSON-файл и преобразует его в JSONL-формат.
func ConvertJSONtoJSONL(filePath string) (io.Reader, string, error) {
	if !strings.HasSuffix(filePath, ".json") {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, "", fmt.Errorf("failed to open file: %v", err)
		}
		return file, "application/octet-stream", nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	decoder := json.NewDecoder(file)
	encoder := json.NewEncoder(&buf)
	for decoder.More() {
		var obj map[string]interface{}
		if err := decoder.Decode(&obj); err != nil {
			return nil, "", fmt.Errorf("failed to decode JSON: %v", err)
		}
		if err := encoder.Encode(obj); err != nil {
			return nil, "", fmt.Errorf("failed to encode JSONL: %v", err)
		}
	}

	return &buf, "application/jsonl", nil
}
