package migrator

import "context"

type Migrator interface {
	Migrate(ctx context.Context) error
	Rollback(ctx context.Context) error
}
