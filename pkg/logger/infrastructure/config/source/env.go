package source

import (
	"context"
	"fmt"
	"os"
	"strconv"

	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
)

type envConfig struct {
	debug bool
}

func NewEnvConfig(_ context.Context) (ctrcfg.Config, error) {
	debugStr := os.Getenv("LOGGER_DEBUG")

	debug, err := strconv.ParseBool(debugStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse logger debug status: %w", err)
	}

	cfg := &envConfig{}
	if err := cfg.SetDebug(debug); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *envConfig) IsDebug() bool {
	return c.debug
}

func (c *envConfig) SetDebug(debug bool) error {
	c.debug = debug
	return nil
}
