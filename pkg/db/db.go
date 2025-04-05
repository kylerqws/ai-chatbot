package db

import (
	"context"
	"fmt"

	"github.com/kylerqws/chatbot/pkg/db/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/db/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/db/infrastructure/migrator"

	ctr "github.com/kylerqws/chatbot/pkg/db/contract"
	ctrcli "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

type db struct {
	client   ctrcli.Client
	migrator ctrmig.Migrator
}

func New(ctx context.Context) (ctr.DB, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("db: failed to load config: %w", err)
	}

	cl := client.New(cfg)
	mig := migrator.New(cl)

	return &db{client: cl, migrator: mig}, nil
}

func (db *db) Client() ctrcli.Client {
	return db.client
}

func (db *db) Migrator() ctrmig.Migrator {
	return db.migrator
}
