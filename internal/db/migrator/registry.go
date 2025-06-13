package migrator

import (
	intmig "github.com/kylerqws/chatbot/internal/db/migrator/migrations"
	bunmig "github.com/uptrace/bun/migrate"
)

// migrationRegistry is the ordered list of schema migration registrations.
var migrationRegistry = []func(*bunmig.Migrations){
	// --- 06.03.2025 ---
	intmig.CreateMessageTable06032025001,
	intmig.CreateServiceTable06032025002,
	intmig.CreateUserTable06032025003,
	intmig.CreateAiChatTable06032025004,
}
