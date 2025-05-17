package client

import (
	"fmt"
	"log"

	"github.com/kylerqws/chatbot/pkg/db/infrastructure/client/dialect"
	"github.com/uptrace/bun"

	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrdlt "github.com/kylerqws/chatbot/pkg/db/contract/client/dialect"
	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"
)

type dbClient struct {
	config  ctrcfg.Config
	dialect ctrdlt.Dialect
	db      *bun.DB
}

func New(cfg ctrcfg.Config) ctrcl.Client {
	return &dbClient{config: cfg}
}

func (c *dbClient) Connect() error {
	dn := c.config.GetDialect()

	switch dn {
	case "sqlite":
		c.dialect = dialect.NewSQLite(c.config)
	default:
		return fmt.Errorf("unsupported database dialect: '%v'", dn)
	}

	err := c.dialect.Connect()
	if err != nil {
		return fmt.Errorf("failed to connect dialect '%v': %w", dn, err)
	}

	c.db = c.dialect.DB()
	return nil
}

func (c *dbClient) Close() error {
	err := c.db.Close()
	if err != nil {
		return fmt.Errorf("failed to close database: %w", err)
	}

	return nil
}

func (c *dbClient) DB() *bun.DB {
	if c.db == nil {
		log.Fatalf("database not initialized")
	}

	return c.db
}
