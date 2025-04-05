package db

import (
	"context"

	"github.com/kylerqws/chatbot/pkg/db/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/db/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/db/infrastructure/migrator"

	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

type DB struct {
	client   ctrcfg.Client
	migrator ctrmig.Migrator
}

func New(ctx context.Context) (*DB, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, err
	}

	cl, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	mig, err := migrator.New(cl)
	if err != nil {
		return nil, err
	}

	return &DB{
		client:   cl,
		migrator: mig,
	}, nil
}

func (db *DB) Client() ctrcfg.Client {
	return db.client
}

func (db *DB) Migrator() ctrmig.Migrator {
	return db.migrator
}
