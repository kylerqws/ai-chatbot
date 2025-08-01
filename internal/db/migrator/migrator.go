package migrator

import (
	"context"
	"fmt"

	"github.com/uptrace/bun/migrate"

	ctr "github.com/kylerqws/chatbot/pkg/db/contract"
	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
)

// Migrate applies all pending database migrations.
func Migrate(ctx context.Context, db ctr.DB) (err error) {
	client := db.Client()
	if err = client.Connect(); err != nil {
		return fmt.Errorf("connect database client: %w", err)
	}
	defer closeClient(client, &err)

	if err = db.Migrator().Migrate(ctx, registeredMigrations()); err != nil {
		return fmt.Errorf("apply migrations: %w", err)
	}
	return err
}

// Rollback rolls back the last applied database migrations.
func Rollback(ctx context.Context, db ctr.DB) (err error) {
	client := db.Client()
	if err = client.Connect(); err != nil {
		return fmt.Errorf("connect database client: %w", err)
	}
	defer closeClient(client, &err)

	if err = db.Migrator().Rollback(ctx, registeredMigrations()); err != nil {
		return fmt.Errorf("rollback migrations: %w", err)
	}
	return err
}

// closeClient closes the database client and updates the error on failure.
func closeClient(cl ctrcl.Client, err *error) {
	if cerr := cl.Close(); cerr != nil && *err == nil {
		*err = fmt.Errorf("close database client: %w", cerr)
	}
}

// registeredMigrations returns the set of migrations created from the registry.
func registeredMigrations() *migrate.Migrations {
	migs := migrate.NewMigrations()
	for i := range registry {
		registry[i](migs)
	}
	return migs
}
