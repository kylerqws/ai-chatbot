package jsonl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// HasJSONSuffix returns true if the file path ends with ".json" (case-insensitive).
func HasJSONSuffix(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), ".json")
}

// ConvertToReader opens a JSON file with an array of objects and returns a JSONL stream.
func ConvertToReader(path string) (io.ReadCloser, error) {
	if !HasJSONSuffix(path) {
		return nil, fmt.Errorf("unsupported file extension (expected .json): %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file '%s': %w", path, err)
	}

	decoder := json.NewDecoder(file)
	tok, err := decoder.Token()
	if err != nil || tok != json.Delim('[') {
		_ = file.Close()
		return nil, fmt.Errorf("expected top-level JSON array: %w", err)
	}

	return startStreamingJSONL(file, decoder)
}

// startStreamingJSONL launches a goroutine that converts JSON array to JSONL lines.
func startStreamingJSONL(file *os.File, decoder *json.Decoder) (io.ReadCloser, error) {
	pr, pw := io.Pipe()

	go func() {
		err := writeJSONLStream(decoder, pw)
		if cerr := file.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close file: %w", cerr)
		}
		_ = pw.CloseWithError(err)
	}()

	return pr, nil
}

// writeJSONLStream reads JSON objects from decoder and writes them as JSONL lines.
func writeJSONLStream(decoder *json.Decoder, w io.Writer) error {
	writer := bufio.NewWriter(w)

	for decoder.More() {
		var obj map[string]any
		if err := decoder.Decode(&obj); err != nil {
			return fmt.Errorf("decode JSON object: %w", err)
		}

		line, err := json.Marshal(obj)
		if err != nil {
			return fmt.Errorf("marshal JSON object: %w", err)
		}

		if _, err := writer.Write(line); err != nil {
			return fmt.Errorf("write JSONL line: %w", err)
		}

		if err := writer.WriteByte('\n'); err != nil {
			return fmt.Errorf("write newline: %w", err)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush writer: %w", err)
	}

	return nil
}
