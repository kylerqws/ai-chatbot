package client

import (
	"fmt"
	"github.com/uptrace/bun"

	"github.com/kylerqws/chatbot/pkg/db/infrastructure/client/dialect"

	ctrclt "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrdlt "github.com/kylerqws/chatbot/pkg/db/contract/client/dialect"
	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"
)

type dbClient struct {
	config  ctrcfg.Config
	dialect ctrdlt.Dialect
	db      *bun.DB
}

func New(cfg ctrcfg.Config) (ctrclt.Client, error) {
	return &dbClient{config: cfg}, nil
}

func (c *dbClient) Connect() error {
	var dlt ctrdlt.Dialect

	dltName := c.config.GetDialect()
	switch dltName {
	case "sqlite":
		dlt = dialect.NewSQLite(c.config)
	default:
		return fmt.Errorf("unsupported databse dialect: %q", dltName)
	}

	c.dialect = dlt
	if err := c.dialect.Connect(); err != nil {
		return err
	}

	return nil
}

func (c *dbClient) Close() error {
	return c.db.Close()
}

func (c *dbClient) DB() *bun.DB {
	return c.db
}
