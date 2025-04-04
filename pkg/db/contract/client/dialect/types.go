package dialect

import (
	"context"
	"database/sql"
)

type Dialect interface {
	Init(ctx context.Context) error
	Connect() (*sql.DB, error)
	Close() error
	DB() *sql.DB
}
