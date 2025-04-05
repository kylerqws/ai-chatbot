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

type envConfig struct {
	baseURL string
	apiKey  string
	timeout time.Duration
}

func NewEnvConfig(_ context.Context) (ctrcfg.Config, error) {
	baseURL := os.Getenv("OPENAI_API_BASE_URL")
	apiKey := os.Getenv("OPENAI_API_KEY")
	timeoutStr := os.Getenv("OPENAI_API_TIMEOUT")

	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		return nil, fmt.Errorf("env config: invalid OPENAI_API_TIMEOUT %q: %w", timeoutStr, err)
	}

	cfg := &envConfig{}
	if err := cfg.SetBaseURL(baseURL); err != nil {
		return nil, fmt.Errorf("env config: %w", err)
	}
	if err := cfg.SetAPIKey(apiKey); err != nil {
		return nil, fmt.Errorf("env config: %w", err)
	}
	if err := cfg.SetTimeout(timeout); err != nil {
		return nil, fmt.Errorf("env config: %w", err)
	}

	return cfg, nil
}

func (c *envConfig) GetBaseURL() string {
	return c.baseURL
}

func (c *envConfig) SetBaseURL(baseURL string) error {
	if strings.TrimSpace(baseURL) == "" {
		return fmt.Errorf("missing required environment variable: OPENAI_API_BASE_URL")
	}
	c.baseURL = baseURL
	return nil
}

func (c *envConfig) GetAPIKey() string {
	return c.apiKey
}

func (c *envConfig) SetAPIKey(apiKey string) error {
	if strings.TrimSpace(apiKey) == "" {
		return fmt.Errorf("missing required environment variable: OPENAI_API_KEY")
	}
	c.apiKey = apiKey
	return nil
}

func (c *envConfig) GetTimeout() time.Duration {
	return c.timeout
}

func (c *envConfig) SetTimeout(seconds int) error {
	if seconds <= 0 {
		return fmt.Errorf("invalid OPENAI_API_TIMEOUT value: %d (must be > 0)", seconds)
	}
	c.timeout = time.Duration(seconds) * time.Second
	return nil
}
