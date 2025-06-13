package contract

import (
	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

// DB defines the internal abstraction over the database client.
type DB interface {
	Client() ctrcl.Client
	Migrator() ctrmig.Migrator
}
