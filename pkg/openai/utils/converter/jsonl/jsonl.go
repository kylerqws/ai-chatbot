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

func ConvertToReader(path string) (reader io.Reader, err error) {
	if !strings.HasSuffix(strings.ToLower(path), ".json") {
		return nil, fmt.Errorf("unsupported file extension (expected .json): %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %v: %w", path, err)
	}

	defer func(file *os.File) {
		if clerr := file.Close(); clerr != nil {
			if err == nil {
				err = fmt.Errorf("failed to close file %v: %w", path, clerr)
			}
		}
	}(file)

	data, err := decodeJSON(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON from %v: %w", path, err)
	}

	jsonl, err := encodeToJSONL(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode to JSONL from %v: %w", path, err)
	}

	reader = bytes.NewReader(jsonl)
	return reader, err
}

func decodeJSON(reader io.Reader) ([]map[string]any, error) {
	var data []map[string]any
	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return data, nil
}

func encodeToJSONL(data []map[string]any) ([]byte, error) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	for _, record := range data {
		line, err := json.Marshal(record)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal record: %w", err)
		}

		_, err = writer.Write(line)
		if err != nil {
			return nil, fmt.Errorf("failed to write line: %w", err)
		}

		err = writer.WriteByte('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to write line break: %w", err)
		}
	}

	err := writer.Flush()
	if err != nil {
		return nil, fmt.Errorf("failed to flush writer: %w", err)
	}

	return buffer.Bytes(), nil
}
