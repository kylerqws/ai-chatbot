package dialect

import "github.com/uptrace/bun"

// Dialect provides access to the database dialect.
type Dialect interface {
	// Connect opens the connection via the dialect.
	Connect() error

	// Close closes the database connection.
	Close() error

	// DB returns the low-level database instance.
	DB() *bun.DB
}
