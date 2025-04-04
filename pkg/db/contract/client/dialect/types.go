package dialect

import "database/sql"

type Dialect interface {
	Connect() error
	Close() error
	DB() *sql.DB
}
