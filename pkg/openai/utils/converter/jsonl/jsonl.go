package jsonl

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func ConvertToReader(path string) (io.Reader, error) {
	if !isJSONFile(path) {
		return nil, fmt.Errorf("file does not have .json extension")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		if clerr := file.Close(); clerr != nil && err == nil {
			err = clerr
		}
	}(file)

	data, err := decodeJSON(file)
	if err != nil {
		return nil, err
	}

	jsonlBytes, err := encodeToJSONL(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonlBytes), nil
}

func decodeJSON(r io.Reader) ([]map[string]any, error) {
	var data []map[string]any
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return data, nil
}

func encodeToJSONL(data []map[string]any) ([]byte, error) {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	for _, record := range data {
		line, err := json.Marshal(record)
		if err != nil {
			return nil, fmt.Errorf("failed to encode JSONL: %w", err)
		}

		if _, err := writer.Write(line); err != nil {
			return nil, err
		}
		if err := writer.WriteByte('\n'); err != nil {
			return nil, err
		}
	}

	err := writer.Flush()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func isJSONFile(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), ".json")
}
