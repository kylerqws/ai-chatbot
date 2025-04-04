package source

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"
)

type envConfig struct {
	dialect string
	dsn     string
	debug   bool
}

func NewEnvConfig(_ context.Context) (ctrcfg.Config, error) {
	dialect := os.Getenv("DB_DIALECT")
	dsn := os.Getenv("DB_DSN")
	debugStr := os.Getenv("DB_DEBUG")

	debug, err := strconv.ParseBool(debugStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database debug status: %w", err)
	}

	cfg := &envConfig{}
	if err := cfg.SetDialect(dialect); err != nil {
		return nil, err
	}
	if err := cfg.SetDSN(dsn); err != nil {
		return nil, err
	}
	if err := cfg.SetDebug(debug); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *envConfig) GetDialect() string {
	return c.dialect
}

func (c *envConfig) SetDialect(dialect string) error {
	if strings.TrimSpace(dialect) == "" {
		return fmt.Errorf("required database dialect is not set")
	}

	c.dialect = dialect
	return nil
}

func (c *envConfig) GetDSN() string {
	return c.dsn
}

func (c *envConfig) SetDSN(dsn string) error {
	if strings.TrimSpace(dsn) == "" {
		return fmt.Errorf("required database source name (DSN) is not set")
	}

	c.dsn = dsn
	return nil
}

func (c *envConfig) IsDebug() bool {
	return c.debug
}

func (c *envConfig) SetDebug(debug bool) error {
	c.debug = debug
	return nil
}
