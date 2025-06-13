package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

// CreateMessageTable06032025001 registers migration for creating the `message` table.
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
			if err != nil {
				return fmt.Errorf("create message table: %w", err)
			}
			return nil
		},
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `DROP TABLE IF EXISTS message;`)
			if err != nil {
				return fmt.Errorf("drop message table: %w", err)
			}
			return nil
		},
	)
}
