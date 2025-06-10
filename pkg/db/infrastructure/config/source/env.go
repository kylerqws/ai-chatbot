package source

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"
)

// envConfig implements the Config interface using environment variables as the source.
type envConfig struct {
	dialect string
	dsn     string
	debug   bool
}

// NewEnvConfig returns a Config implementation populated from environment variables.
// Falls back to default values when optional variables are missing.
func NewEnvConfig(_ context.Context) (ctrcfg.Config, error) {
	dialect := os.Getenv("DB_DIALECT")
	dsn := os.Getenv("DB_DSN")
	debugStr := os.Getenv("DB_DEBUG")

	debug, err := strconv.ParseBool(debugStr)
	if err != nil {
		debug = DefaultDebug
	}

	cfg := &envConfig{}
	if err := cfg.SetDialect(dialect); err != nil {
		return nil, fmt.Errorf("set dialect: %w", err)
	}
	if err := cfg.SetDSN(dsn); err != nil {
		return nil, fmt.Errorf("set DSN: %w", err)
	}
	if err := cfg.SetDebug(debug); err != nil {
		return nil, fmt.Errorf("set debug: %w", err)
	}

	return cfg, nil
}

// GetDialect returns the database dialect.
func (c *envConfig) GetDialect() string {
	return c.dialect
}

// SetDialect sets the database dialect.
// If the input is empty or whitespace, a default value is used.
func (c *envConfig) SetDialect(dialect string) error {
	if dialect = strings.TrimSpace(dialect); dialect == "" {
		dialect = DefaultDialect
	}
	c.dialect = dialect
	return nil
}

// GetDSN returns the database connection string.
func (c *envConfig) GetDSN() string {
	return c.dsn
}

// SetDSN sets the database connection string (DSN).
// If the input is empty or whitespace, a default value is used.
func (c *envConfig) SetDSN(dsn string) error {
	if dsn = strings.TrimSpace(dsn); dsn == "" {
		dsn = DefaultDSN
	}
	c.dsn = dsn
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
