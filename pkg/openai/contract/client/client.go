package client

import (
	"context"
	"io"
)

// Client defines the interface for sending HTTP requests to the OpenAI API.
// It supports JSON-encoded, multipart/form-data, and raw requests.
type Client interface {
	// RequestJSON sends an HTTP request with a JSON-encoded body.
	// Returns the raw response body or an error.
	RequestJSON(ctx context.Context, method, path string, body any) ([]byte, error)

	// RequestMultipart sends a multipart/form-data POST request to the specified path.
	// The body map must contain a "file" entry with the path to the file.
	RequestMultipart(ctx context.Context, path string, body map[string]string) ([]byte, error)

	// RequestRaw sends an HTTP request with the given method and optional body.
	// Suitable for GET, DELETE, or other custom requests.
	RequestRaw(ctx context.Context, method, path string, body io.Reader) ([]byte, error)
}
