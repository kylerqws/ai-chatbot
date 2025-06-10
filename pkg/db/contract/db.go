package contract

import (
	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

// DB aggregates access to database functionality.
type DB interface {
	// Client returns the database client instance.
	Client() ctrcl.Client

	// Migrator returns the database migration manager.
	Migrator() ctrmig.Migrator
}
