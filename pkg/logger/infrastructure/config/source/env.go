package source

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
)

type envConfig struct {
	writer string
	debug  bool
}

func NewEnvConfig(_ context.Context) (ctrcfg.Config, error) {
	writer := os.Getenv("LOGGER_WRITER")
	debugStr := os.Getenv("LOGGER_DEBUG")

	debug, err := strconv.ParseBool(debugStr)
	if err != nil {
		return nil, fmt.Errorf("invalid LOGGER_DEBUG: %w", err)
	}

	cfg := &envConfig{}
	if err := cfg.SetWriter(writer); err != nil {
		return nil, err
	}
	if err := cfg.SetDebug(debug); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *envConfig) GetWriter() string {
	return c.writer
}

func (c *envConfig) SetWriter(writer string) error {
	if strings.TrimSpace(writer) == "" {
		return fmt.Errorf("missing LOGGER_WRITER")
	}

	c.writer = writer
	return nil
}

func (c *envConfig) IsDebug() bool {
	return c.debug
}

func (c *envConfig) SetDebug(debug bool) error {
	c.debug = debug
	return nil
}
