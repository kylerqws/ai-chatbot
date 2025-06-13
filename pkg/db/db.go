package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/kylerqws/chatbot/pkg/db/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/db/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/db/infrastructure/migrator"

	ctr "github.com/kylerqws/chatbot/pkg/db/contract"
	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

// entrypoint aggregates access to the database client and migrator.
type entrypoint struct {
	ctx context.Context
	cfg ctrcfg.Config

	cl     ctrcl.Client
	clOnce sync.Once

	mig     ctrmig.Migrator
	migOnce sync.Once
}

// New creates and returns a new DB access object.
func New(ctx context.Context) (ctr.DB, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("load db config: %w", err)
	}
	return &entrypoint{ctx: ctx, cfg: cfg}, nil
}

// Client returns the initialized database client.
func (e *entrypoint) Client() ctrcl.Client {
	e.clOnce.Do(func() {
		e.cl = client.New(e.cfg)
	})
	return e.cl
}

// Migrator returns the initialized database migrator.
func (e *entrypoint) Migrator() ctrmig.Migrator {
	e.migOnce.Do(func() {
		e.mig = migrator.New(e.Client())
	})
	return e.mig
}
