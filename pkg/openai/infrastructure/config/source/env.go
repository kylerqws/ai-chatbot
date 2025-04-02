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
		return nil, fmt.Errorf("failed to parse OpenAI timeout value: %w", err)
	}

	cfg := &envConfig{}
	if err := cfg.SetBaseURL(baseURL); err != nil {
		return nil, err
	}
	if err := cfg.SetAPIKey(apiKey); err != nil {
		return nil, err
	}
	if err := cfg.SetTimeout(timeout); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *envConfig) GetBaseURL() string {
	return c.baseURL
}

func (c *envConfig) SetBaseURL(baseURL string) error {
	if strings.TrimSpace(baseURL) == "" {
		return fmt.Errorf("required OpenAI base URL is not set")
	}

	c.baseURL = baseURL
	return nil
}

func (c *envConfig) GetAPIKey() string {
	return c.apiKey
}

func (c *envConfig) SetAPIKey(apiKey string) error {
	if strings.TrimSpace(apiKey) == "" {
		return fmt.Errorf("required OpenAI key is not set")
	}

	c.apiKey = apiKey
	return nil
}

func (c *envConfig) GetTimeout() time.Duration {
	return c.timeout
}

func (c *envConfig) SetTimeout(seconds int) error {
	if seconds <= 0 {
		return fmt.Errorf("invalid value for OpenAI timeout: %q", seconds)
	}

	c.timeout = time.Duration(seconds) * time.Second
	return nil
}
