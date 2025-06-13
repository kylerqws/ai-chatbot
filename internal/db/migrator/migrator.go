package migrator

import (
	"context"
	"fmt"

	"github.com/uptrace/bun/migrate"

	ctr "github.com/kylerqws/chatbot/pkg/db/contract"
	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
)

// Migrate connects to the database and applies all pending migrations.
func Migrate(ctx context.Context, db ctr.DB) (err error) {
	client := db.Client()
	if err = client.Connect(); err != nil {
		return fmt.Errorf("connect database client: %w", err)
	}
	defer closeClient(client, &err)

	if err = db.Migrator().Migrate(ctx, registeredMigrations()); err != nil {
		return fmt.Errorf("apply migrations: %w", err)
	}
	return nil
}

// Rollback connects to the database and rolls back the last applied migration.
func Rollback(ctx context.Context, db ctr.DB) (err error) {
	client := db.Client()
	if err = client.Connect(); err != nil {
		return fmt.Errorf("connect database client: %w", err)
	}
	defer closeClient(client, &err)

	if err = db.Migrator().Rollback(ctx, registeredMigrations()); err != nil {
		return fmt.Errorf("rollback migrations: %w", err)
	}
	return nil
}

// closeClient ensures the database client is closed and captures close errors.
func closeClient(cl ctrcl.Client, err *error) {
	if cerr := cl.Close(); cerr != nil && *err == nil {
		*err = fmt.Errorf("close database client: %w", cerr)
	}
}

// registeredMigrations returns a Migrations object populated from the registry.
func registeredMigrations() *migrate.Migrations {
	migs := migrate.NewMigrations()
	for _, register := range migrationRegistry {
		register(migs)
	}
	return migs
}
