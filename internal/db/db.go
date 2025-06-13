package db

import (
	"context"
	"fmt"
	"log"

	"github.com/kylerqws/chatbot/pkg/db"
	"github.com/uptrace/bun"

	ctrint "github.com/kylerqws/chatbot/internal/db/contract"
	ctrpkg "github.com/kylerqws/chatbot/pkg/db/contract"
)

// proxy is an internal wrapper over the base DB client.
type proxy struct {
	db ctrpkg.DB
}

// New creates a new internal DB proxy over the base DB client.
func New(ctx context.Context) ctrint.DB {
	instance, err := db.New(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("create base DB client: %w", err))
	}
	return &proxy{db: instance}
}

// Connect opens the database connection.
func (p *proxy) Connect() error {
	return p.db.Client().Connect()
}

// Close closes the database connection.
func (p *proxy) Close() error {
	return p.db.Client().Close()
}

// DB returns the low-level database instance.
func (p *proxy) DB() *bun.DB {
	return p.db.Client().DB()
}
