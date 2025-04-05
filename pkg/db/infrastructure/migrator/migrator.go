package migrator

import (
	"context"
	"fmt"

	"github.com/uptrace/bun/migrate"

	ctrcli "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

type migrator struct {
	client ctrcli.Client
}

func New(cl ctrcli.Client) (ctrmig.Migrator, error) {
	return &migrator{client: cl}, nil
}

func (m *migrator) Migrate(ctx context.Context, mgs *migrate.Migrations) error {
	migrator, err := m.newMigrator(ctx, mgs)
	if err != nil {
		return err
	}

	_, err = migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}

func (m *migrator) Rollback(ctx context.Context, mgs *migrate.Migrations) error {
	migrator, err := m.newMigrator(ctx, mgs)
	if err != nil {
		return err
	}

	_, err = migrator.Rollback(ctx)
	if err != nil {
		return fmt.Errorf("rollback failed: %w", err)
	}

	return nil
}

func (m *migrator) newMigrator(ctx context.Context, mgs *migrate.Migrations) (*migrate.Migrator, error) {
	migrator := migrate.NewMigrator(m.client.DB(), mgs)
	if err := migrator.Init(ctx); err != nil {
		return nil, fmt.Errorf("failed to init migrator: %w", err)
	}

	return migrator, nil
}
