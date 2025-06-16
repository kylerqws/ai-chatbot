package migrator

import (
	"context"
	"fmt"

	"github.com/uptrace/bun/migrate"

	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

// migrator handles database schema migrations.
type migrator struct {
	client ctrcl.Client
}

// New returns a new migrator using the given database client.
func New(cl ctrcl.Client) ctrmig.Migrator {
	return &migrator{client: cl}
}

// Migrate applies all pending migrations.
func (m *migrator) Migrate(ctx context.Context, mgs *migrate.Migrations) error {
	mig, err := m.newMigrator(ctx, mgs)
	if err != nil {
		return fmt.Errorf("create migrator: %w", err)
	}
	if _, err := mig.Migrate(ctx); err != nil {
		return fmt.Errorf("apply migrations: %w", err)
	}
	return nil
}

// Rollback rolls back the last applied migrations.
func (m *migrator) Rollback(ctx context.Context, mgs *migrate.Migrations) error {
	mig, err := m.newMigrator(ctx, mgs)
	if err != nil {
		return fmt.Errorf("create migrator: %w", err)
	}
	if _, err := mig.Rollback(ctx); err != nil {
		return fmt.Errorf("rollback migrations: %w", err)
	}
	return nil
}

// newMigrator creates and initializes a migrator with the provided migrations.
func (m *migrator) newMigrator(ctx context.Context, mgs *migrate.Migrations) (*migrate.Migrator, error) {
	mig := migrate.NewMigrator(m.client.DB(), mgs)
	if err := mig.Init(ctx); err != nil {
		return nil, fmt.Errorf("init migrator: %w", err)
	}
	return mig, nil
}
