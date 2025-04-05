package migration

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func CreateMessageTable06032025001(migrations *migrate.Migrations) {
	migrations.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `
				CREATE TABLE IF NOT EXISTS message (
					entity_id INTEGER PRIMARY KEY AUTOINCREMENT,
					request TEXT NOT NULL,
					response TEXT NOT NULL,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				);
			`)

			return err
		},
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `DROP TABLE IF EXISTS message;`)

			return err
		},
	)
}
