package db

import (
	"context"
	"fmt"

	"github.com/kylerqws/chatbot/pkg/db/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/db/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/db/infrastructure/migrator"

	ctr "github.com/kylerqws/chatbot/pkg/db/contract"
	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

type db struct {
	client   ctrcl.Client
	migrator ctrmig.Migrator
}

func New(ctx context.Context) (ctr.DB, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	cl := client.New(cfg)
	mig := migrator.New(cl)

	return &db{client: cl, migrator: mig}, nil
}

func (db *db) Client() ctrcl.Client {
	return db.client
}

func (db *db) Migrator() ctrmig.Migrator {
	return db.migrator
}
