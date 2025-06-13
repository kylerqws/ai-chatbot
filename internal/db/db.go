package db

import (
	"context"
	"fmt"
	"log"

	"github.com/kylerqws/chatbot/pkg/db"

	ctrint "github.com/kylerqws/chatbot/internal/db/contract"
	ctrpkg "github.com/kylerqws/chatbot/pkg/db/contract"

	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
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

func (p *proxy) Client() ctrcl.Client {
	return p.db.Client()
}

func (p *proxy) Migrator() ctrmig.Migrator {
	return p.db.Migrator()
}
