package migrator

import (
	"context"
	"github.com/uptrace/bun/migrate"
)

type Migrator interface {
	Migrate(ctx context.Context, mgs *migrate.Migrations) error
	Rollback(ctx context.Context, mgs *migrate.Migrations) error
}
