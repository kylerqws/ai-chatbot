package migrator

import (
	"context"
	"fmt"
	"github.com/uptrace/bun/migrate"

	ctrcli "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

type migrator struct {
	client     ctrcli.Client
	migrations *migrate.Migrations
}

func New(cl ctrcli.Client) (ctrmig.Migrator, error) {
	return &migrator{client: cl}, nil
}

func (m *migrator) Migrate(ctx context.Context) error {
	migrator, err := m.newMigrator(ctx)
	if err != nil {
		return err
	}

	_, err = migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}

func (m *migrator) Rollback(ctx context.Context) error {
	migrator, err := m.newMigrator(ctx)
	if err != nil {
		return err
	}

	_, err = migrator.Rollback(ctx)
	if err != nil {
		return fmt.Errorf("rollback migrations failed: %w", err)
	}

	return nil
}

func (m *migrator) newMigrator(ctx context.Context) (*migrate.Migrator, error) {
	migrator := migrate.NewMigrator(m.client.DB(), m.migrations)
	if err := migrator.Init(ctx); err != nil {
		return nil, fmt.Errorf("failed to init migrations: %w", err)
	}

	return migrator, nil
}
