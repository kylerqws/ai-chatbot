package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

// CreateUserTable06032025003 registers migration for creating the `user` table.
func CreateUserTable06032025003(migrations *migrate.Migrations) {
	migrations.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `
				CREATE TABLE IF NOT EXISTS user (
					entity_id INTEGER PRIMARY KEY AUTOINCREMENT,
					service_id INTEGER NOT NULL,
					user_id TEXT NOT NULL,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					UNIQUE(service_id, user_id),
					FOREIGN KEY (service_id) REFERENCES service(entity_id) ON DELETE CASCADE
				);
			`)
			if err != nil {
				return fmt.Errorf("create `user` table: %w", err)
			}
			return nil
		},
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `DROP TABLE IF EXISTS user;`)
			if err != nil {
				return fmt.Errorf("drop `user` table: %w", err)
			}
			return nil
		},
	)
}
