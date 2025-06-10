package client

import "github.com/uptrace/bun"

// Client provides access to the database connection.
type Client interface {
	// Connect opens the database connection.
	Connect() error

	// Close closes the database connection.
	Close() error

	// DB returns the low-level database instance.
	DB() *bun.DB
}
