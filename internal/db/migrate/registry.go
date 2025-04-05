package migrate

import (
	"github.com/kylerqws/chatbot/internal/db/migrate/migrations"
	"github.com/uptrace/bun/migrate"
)

func Migrations() *migrate.Migrations {
	migs := migrate.NewMigrations()

	migrations.CreateMessageTable06032025001(migs)
	migrations.CreateServiceTable06032025002(migs)
	migrations.CreateUserTable06032025003(migs)
	migrations.CreateAiChatTable06032025004(migs)

	return migs
}
