package migration

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

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

			return err
		},
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `DROP TABLE IF EXISTS user;`)

			return err
		},
	)
}
