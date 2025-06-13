package source

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
)

// envConfig implements the Config interface using environment variables as the source.
type envConfig struct {
	baseURL string
	apiKey  string

	timeout     time.Duration
	tlsTimeout  time.Duration
	respTimeout time.Duration
}

// NewEnvConfig returns a Config implementation populated from environment variables.
// Falls back to default values when optional variables are missing.
func NewEnvConfig(_ context.Context) (ctrcfg.Config, error) {
	url := os.Getenv("OPENAI_API_BASE_URL")
	key := os.Getenv("OPENAI_API_KEY")

	timeoutStr := os.Getenv("OPENAI_API_TIMEOUT")
	tlsTimeoutStr := os.Getenv("OPENAI_API_TLS_HANDSHAKE_TIMEOUT")
	respTimeoutStr := os.Getenv("OPENAI_API_RESPONSE_HEADER_TIMEOUT")

	timeout, err := strconv.ParseUint(timeoutStr, 10, 64)
	if err != nil {
		timeout = DefaultTimeout
	}
	tlsTimeout, err := strconv.ParseUint(tlsTimeoutStr, 10, 64)
	if err != nil {
		tlsTimeout = DefaultTLSHandshakeTimeout
	}
	respTimeout, err := strconv.ParseUint(respTimeoutStr, 10, 64)
	if err != nil {
		respTimeout = DefaultResponseHeaderTimeout
	}

	cfg := &envConfig{}
	if err := cfg.SetBaseURL(url); err != nil {
		return nil, fmt.Errorf("set base URL: %w", err)
	}
	if err := cfg.SetAPIKey(key); err != nil {
		return nil, fmt.Errorf("set API key: %w", err)
	}
	if err := cfg.SetTimeout(timeout); err != nil {
		return nil, fmt.Errorf("set timeout: %w", err)
	}
	if err := cfg.SetTLSHandshakeTimeout(tlsTimeout); err != nil {
		return nil, fmt.Errorf("set TLS handshake timeout: %w", err)
	}
	if err := cfg.SetResponseHeaderTimeout(respTimeout); err != nil {
		return nil, fmt.Errorf("set response header timeout: %w", err)
	}

	return cfg, nil
}

// GetBaseURL returns the OpenAI API base URL.
func (c *envConfig) GetBaseURL() string {
	return c.baseURL
}

// SetBaseURL sets the OpenAI API base URL.
// If the input is empty or whitespace, a default value is used.
func (c *envConfig) SetBaseURL(baseURL string) error {
	if baseURL = strings.TrimSpace(baseURL); baseURL == "" {
		baseURL = DefaultBaseURL
	}
	c.baseURL = baseURL
	return nil
}

// GetAPIKey returns the API key used for authentication.
func (c *envConfig) GetAPIKey() string {
	return c.apiKey
}

// SetAPIKey sets the API key for authentication.
// Returns an error if the key is empty or only whitespace.
func (c *envConfig) SetAPIKey(apiKey string) error {
	if apiKey = strings.TrimSpace(apiKey); apiKey == "" {
		return fmt.Errorf("missing required OpenAI API key")
	}
	c.apiKey = apiKey
	return nil
}

// GetTimeout returns the HTTP client timeout duration.
func (c *envConfig) GetTimeout() time.Duration {
	return c.timeout
}

// SetTimeout sets the HTTP client timeout in seconds.
// Returns a default value if the provided timeout is zero.
func (c *envConfig) SetTimeout(seconds uint64) error {
	if seconds == 0 {
		seconds = DefaultTimeout
	}
	c.timeout = time.Duration(seconds) * time.Second
	return nil
}

// GetTLSHandshakeTimeout returns the TLS handshake timeout duration.
func (c *envConfig) GetTLSHandshakeTimeout() time.Duration {
	return c.tlsTimeout
}

// SetTLSHandshakeTimeout sets the TLS handshake timeout in seconds.
// Returns a default value if the provided timeout is zero.
func (c *envConfig) SetTLSHandshakeTimeout(seconds uint64) error {
	if seconds == 0 {
		seconds = DefaultTLSHandshakeTimeout
	}
	c.tlsTimeout = time.Duration(seconds) * time.Second
	return nil
}

// GetResponseHeaderTimeout returns the response header timeout duration.
func (c *envConfig) GetResponseHeaderTimeout() time.Duration {
	return c.respTimeout
}

// SetResponseHeaderTimeout sets the response header timeout in seconds.
// Returns a default value if the provided timeout is zero.
func (c *envConfig) SetResponseHeaderTimeout(seconds uint64) error {
	if seconds == 0 {
		seconds = DefaultResponseHeaderTimeout
	}
	c.respTimeout = time.Duration(seconds) * time.Second
	return nil
}
