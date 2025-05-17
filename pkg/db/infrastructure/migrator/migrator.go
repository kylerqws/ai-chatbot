package migrator

import (
	"context"
	"fmt"

	"github.com/uptrace/bun/migrate"

	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

type migrator struct {
	client ctrcl.Client
}

func New(cl ctrcl.Client) ctrmig.Migrator {
	return &migrator{client: cl}
}

func (m *migrator) Migrate(ctx context.Context, mgs *migrate.Migrations) error {
	mig, err := m.newMigrator(ctx, mgs)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}

	_, err = mig.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("failed to migrate migrations: %w", err)
	}

	return nil
}

func (m *migrator) Rollback(ctx context.Context, mgs *migrate.Migrations) error {
	mig, err := m.newMigrator(ctx, mgs)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}

	_, err = mig.Rollback(ctx)
	if err != nil {
		return fmt.Errorf("failed to rollback migrations: %w", err)
	}

	return nil
}

func (m *migrator) newMigrator(ctx context.Context, mgs *migrate.Migrations) (*migrate.Migrator, error) {
	mig := migrate.NewMigrator(m.client.DB(), mgs)

	err := mig.Init(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to init migrator: %w", err)
	}

	return mig, nil
}
