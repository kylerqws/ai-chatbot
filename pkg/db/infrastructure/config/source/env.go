package source

import (
	"context"
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
		debug = DefaultDebug
	}

	cfg := &envConfig{}
	if err = cfg.SetDialect(dialect); err != nil {
		return nil, err
	}
	if err = cfg.SetDSN(dsn); err != nil {
		return nil, err
	}
	if err = cfg.SetDebug(debug); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *envConfig) GetDialect() string {
	return c.dialect
}

func (c *envConfig) SetDialect(dialect string) error {
	dialect = strings.TrimSpace(dialect)
	if dialect == "" {
		dialect = DefaultDialect
	}

	c.dialect = dialect
	return nil
}

func (c *envConfig) GetDSN() string {
	return c.dsn
}

func (c *envConfig) SetDSN(dsn string) error {
	dsn = strings.TrimSpace(dsn)
	if dsn == "" {
		dsn = DefaultDsn
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
