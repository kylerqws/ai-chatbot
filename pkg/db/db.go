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

// db provides access to the database client and migrator.
type db struct {
	ctx context.Context
	cfg ctrcfg.Config

	client ctrcl.Client
	clOnce sync.Once

	migrator ctrmig.Migrator
	migOnce  sync.Once
}

// New creates and returns a new DB access object.
func New(ctx context.Context) (ctr.DB, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("load db config: %w", err)
	}
	return &db{ctx: ctx, cfg: cfg}, nil
}

// Client returns the initialized database client.
func (d *db) Client() ctrcl.Client {
	d.clOnce.Do(func() {
		d.client = client.New(d.cfg)
	})
	return d.client
}

// Migrator returns the initialized database migrator.
func (d *db) Migrator() ctrmig.Migrator {
	d.migOnce.Do(func() {
		d.migrator = migrator.New(d.Client())
	})
	return d.migrator
}
