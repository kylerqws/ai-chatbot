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
		return nil, fmt.Errorf("jsonl: unsupported file extension (expected .json): %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("jsonl: failed to open file %q: %w", path, err)
	}
	defer func(file *os.File) {
		if clerr := file.Close(); clerr != nil && err == nil {
			err = fmt.Errorf("jsonl: failed to close file %q: %w", path, clerr)
		}
	}(file)

	data, err := decodeJSON(file)
	if err != nil {
		return nil, fmt.Errorf("jsonl: failed to decode JSON from %q: %w", path, err)
	}

	jsonlBytes, err := encodeToJSONL(data)
	if err != nil {
		return nil, fmt.Errorf("jsonl: failed to encode to JSONL from %q: %w", path, err)
	}

	return bytes.NewReader(jsonlBytes), nil
}

func decodeJSON(r io.Reader) ([]map[string]any, error) {
	var data []map[string]any
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("jsonl: failed to decode JSON: %w", err)
	}

	return data, nil
}

func encodeToJSONL(data []map[string]any) ([]byte, error) {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	for _, record := range data {
		line, err := json.Marshal(record)
		if err != nil {
			return nil, fmt.Errorf("jsonl: failed to marshal record: %w", err)
		}

		if _, err := writer.Write(line); err != nil {
			return nil, fmt.Errorf("jsonl: failed to write line: %w", err)
		}
		if err := writer.WriteByte('\n'); err != nil {
			return nil, fmt.Errorf("jsonl: failed to write line break: %w", err)
		}
	}

	if err := writer.Flush(); err != nil {
		return nil, fmt.Errorf("jsonl: failed to flush writer: %w", err)
	}

	return buf.Bytes(), nil
}

func isJSONFile(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), ".json")
}
