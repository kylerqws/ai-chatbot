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
	timeout time.Duration
}

// NewEnvConfig returns a Config implementation populated from environment variables.
// Falls back to default values when optional variables are missing.
func NewEnvConfig(_ context.Context) (ctrcfg.Config, error) {
	baseURL := os.Getenv("OPENAI_API_BASE_URL")
	apiKey := os.Getenv("OPENAI_API_KEY")
	timeoutStr := os.Getenv("OPENAI_API_TIMEOUT")

	timeout, err := strconv.ParseUint(timeoutStr, 10, 64)
	if err != nil {
		timeout = DefaultTimeout
	}

	cfg := &envConfig{}
	if err := cfg.SetBaseURL(baseURL); err != nil {
		return nil, fmt.Errorf("set base URL: %w", err)
	}
	if err := cfg.SetAPIKey(apiKey); err != nil {
		return nil, fmt.Errorf("set API key: %w", err)
	}
	if err := cfg.SetTimeout(timeout); err != nil {
		return nil, fmt.Errorf("set timeout: %w", err)
	}

	return cfg, nil
}

// GetBaseURL returns the configured OpenAI API base URL.
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

// GetAPIKey returns the configured OpenAI API key.
func (c *envConfig) GetAPIKey() string {
	return c.apiKey
}

// SetAPIKey sets the OpenAI API key.
// Returns an error if the key is empty or only whitespace.
func (c *envConfig) SetAPIKey(apiKey string) error {
	if apiKey = strings.TrimSpace(apiKey); apiKey == "" {
		return fmt.Errorf("missing required OpenAI API key")
	}
	c.apiKey = apiKey
	return nil
}

// GetTimeout returns the configured HTTP client timeout.
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
