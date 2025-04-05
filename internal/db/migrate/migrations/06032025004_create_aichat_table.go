package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

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

			return err
		},
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `DROP TABLE IF EXISTS aichat;`)

			return err
		},
	)
}
