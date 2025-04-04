package dialect

import "github.com/uptrace/bun"

type Dialect interface {
	Connect() error
	Close() error
	DB() *bun.DB
}
