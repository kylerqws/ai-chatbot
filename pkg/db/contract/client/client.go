package client

import (
	"context"
	"database/sql"
)

type Client interface {
	Init(context.Context) error
	Connect() (*sql.DB, error)
	Close() error
	DB() *sql.DB
}
