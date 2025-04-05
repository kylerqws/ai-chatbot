package contract

import (
	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

type DB interface {
	Client() ctrcfg.Client
	Migrator() ctrmig.Migrator
}
