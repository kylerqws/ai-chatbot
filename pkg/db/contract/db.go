package contract

import (
	ctrcli "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

type DB interface {
	Client() ctrcli.Client
	Migrator() ctrmig.Migrator
}
