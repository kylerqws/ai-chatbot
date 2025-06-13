package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

// CreateAiChatTable06032025004 registers migration for creating the `aichat` table.
func CreateAiChatTable06032025004(migrations *migrate.Migrations) {
	migrations.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `
				CREATE TABLE IF NOT EXISTS aichat (
					entity_id INTEGER PRIMARY KEY AUTOINCREMENT,
					user_id INTEGER NOT NULL,
					message_id INTEGER NOT NULL,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					FOREIGN KEY (user_id) REFERENCES user(entity_id) ON DELETE CASCADE,
					FOREIGN KEY (message_id) REFERENCES message(entity_id) ON DELETE CASCADE
				);
			`)
			if err != nil {
				return fmt.Errorf("create aichat table: %w", err)
			}
			return nil
		},
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `DROP TABLE IF EXISTS aichat;`)
			if err != nil {
				return fmt.Errorf("drop aichat table: %w", err)
			}
			return nil
		},
	)
}
