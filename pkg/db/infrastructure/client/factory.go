package client

import (
	"fmt"

	"github.com/kylerqws/chatbot/pkg/db/infrastructure/client/dialect"
	"github.com/uptrace/bun"

	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrdlt "github.com/kylerqws/chatbot/pkg/db/contract/client/dialect"
	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"
)

// dbClient implements the Client interface using a specific SQL dialect.
type dbClient struct {
	config  ctrcfg.Config
	dialect ctrdlt.Dialect
	db      *bun.DB
}

// New returns a new database client.
func New(cfg ctrcfg.Config) ctrcl.Client {
	return &dbClient{config: cfg}
}

// Connect opens the database connection via the configured dialect.
func (c *dbClient) Connect() error {
	dn := c.config.GetDialect()

	switch dn {
	case "sqlite":
		c.dialect = dialect.NewSQLite(c.config)
	default:
		return fmt.Errorf("unsupported database dialect: '%s'", dn)
	}

	if err := c.dialect.Connect(); err != nil {
		return fmt.Errorf("connect with dialect '%s': %w", dn, err)
	}

	c.db = c.dialect.DB()
	return nil
}

// Close closes the database connection.
func (c *dbClient) Close() error {
	if err := c.db.Close(); err != nil {
		return fmt.Errorf("close database: %w", err)
	}
	return nil
}

// DB returns the low-level database instance.
func (c *dbClient) DB() *bun.DB {
	if c.db == nil {
		panic("database not initialized")
	}
	return c.db
}
