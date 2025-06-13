package migrator

import (
	"context"
	"github.com/uptrace/bun/migrate"
)

// Migrator defines methods to apply and rollback database migrations.
type Migrator interface {
	// Migrate applies all pending database migrations.
	Migrate(ctx context.Context, mgs *migrate.Migrations) error

	// Rollback rolls back the last applied database migrations.
	Rollback(ctx context.Context, mgs *migrate.Migrations) error
}
