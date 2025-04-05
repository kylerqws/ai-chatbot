package client

import (
	"fmt"

	"github.com/kylerqws/chatbot/pkg/db/infrastructure/client/dialect"
	"github.com/uptrace/bun"

	ctrcli "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrdlt "github.com/kylerqws/chatbot/pkg/db/contract/client/dialect"
	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"
)

type dbClient struct {
	config  ctrcfg.Config
	dialect ctrdlt.Dialect
	db      *bun.DB
}

func New(cfg ctrcfg.Config) ctrcli.Client {
	return &dbClient{config: cfg}
}

func (c *dbClient) Connect() error {
	dltName := c.config.GetDialect()

	var dlt ctrdlt.Dialect
	switch dltName {
	case "sqlite":
		dlt = dialect.NewSQLite(c.config)
	default:
		return fmt.Errorf("client: unsupported database dialect: %q", dltName)
	}

	c.dialect = dlt
	if err := c.dialect.Connect(); err != nil {
		return fmt.Errorf("client: failed to connect dialect %q: %w", dltName, err)
	}

	return nil
}

func (c *dbClient) Close() error {
	if err := c.db.Close(); err != nil {
		return fmt.Errorf("client: failed to close database: %w", err)
	}

	return nil
}

func (c *dbClient) DB() *bun.DB {
	return c.db
}
