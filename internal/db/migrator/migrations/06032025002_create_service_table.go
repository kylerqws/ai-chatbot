package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func CreateServiceTable06032025002(migrations *migrate.Migrations) {
	migrations.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `
				CREATE TABLE IF NOT EXISTS service (
					entity_id INTEGER PRIMARY KEY AUTOINCREMENT,
					code TEXT UNIQUE NOT NULL,
					name TEXT UNIQUE NOT NULL,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				);
			`)

			return err
		},
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `DROP TABLE IF EXISTS service;`)

			return err
		},
	)
}
