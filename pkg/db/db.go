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

// manager provides access to the database client and migrator.
type manager struct {
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
	return &manager{ctx: ctx, cfg: cfg}, nil
}

// Client returns the initialized database client.
func (m *manager) Client() ctrcl.Client {
	m.clOnce.Do(func() {
		m.cl = client.New(m.cfg)
	})
	return m.cl
}

// Migrator returns the initialized database migrator.
func (m *manager) Migrator() ctrmig.Migrator {
	m.migOnce.Do(func() {
		m.mig = migrator.New(m.Client())
	})
	return m.mig
}
