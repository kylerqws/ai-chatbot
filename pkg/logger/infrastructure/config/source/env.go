package source

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
)

// envConfig implements the Config interface using environment variables as the source.
type envConfig struct {
	writer string
	debug  bool
}

// NewEnvConfig returns a Config implementation populated from environment variables.
// Falls back to default values when optional variables are missing.
func NewEnvConfig(_ context.Context) (ctrcfg.Config, error) {
	writer := os.Getenv("LOGGER_WRITER")
	debugStr := os.Getenv("LOGGER_DEBUG")

	debug, err := strconv.ParseBool(debugStr)
	if err != nil {
		debug = DefaultDebug
	}

	cfg := &envConfig{}
	if err := cfg.SetWriter(writer); err != nil {
		return nil, fmt.Errorf("set writer: %w", err)
	}
	if err := cfg.SetDebug(debug); err != nil {
		return nil, fmt.Errorf("set debug: %w", err)
	}

	return cfg, nil
}

// GetWriter returns the log writer type.
func (c *envConfig) GetWriter() string {
	return c.writer
}

// SetWriter sets the log writer type.
// If the input is empty or whitespace, a default value is used.
func (c *envConfig) SetWriter(writer string) error {
	if writer = strings.TrimSpace(writer); writer == "" {
		writer = DefaultWriter
	}
	c.writer = writer
	return nil
}

// IsDebug returns whether debug mode is enabled.
func (c *envConfig) IsDebug() bool {
	return c.debug
}

// SetDebug enables or disables debug mode.
// This method always succeeds since the input is a valid boolean.
func (c *envConfig) SetDebug(debug bool) error {
	c.debug = debug
	return nil
}
