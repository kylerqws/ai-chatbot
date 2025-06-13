package contract

import ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"

// DB defines the internal abstraction over the database client.
type DB interface {
	ctrcl.Client
}
