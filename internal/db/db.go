package db

import (
	"context"

	"github.com/kylerqws/chatbot/internal/db/migration"
	"github.com/kylerqws/chatbot/pkg/db"
	ctr "github.com/kylerqws/chatbot/pkg/db/contract"
	"github.com/uptrace/bun/migrate"
)

func New(ctx context.Context) (ctr.DB, error) {
	return db.New(ctx)
}

func Migrations() *migrate.Migrations {
	migrations := migrate.NewMigrations()

	migration.CreateMessageTable06032025001(migrations)
	migration.CreateServiceTable06032025002(migrations)
	migration.CreateUserTable06032025003(migrations)
	migration.CreateAiChatTable06032025004(migrations)

	return migrations
}
