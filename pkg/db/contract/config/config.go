package config

type (
	// SourceType defines the type of configuration source (e.g., env, file, etc.).
	SourceType string

	// ContextKey is a custom type to avoid key collisions in context.
	ContextKey string
)

const (
	// EnvSourceType uses environment variables as the config source.
	EnvSourceType SourceType = "env"

	// SourceTypeKey is the context key used to define the configuration source type.
	SourceTypeKey ContextKey = "sourceType"

	// DefaultSourceType is the fallback configuration source type used when none is provided.
	DefaultSourceType = EnvSourceType
)

// Config defines database configuration settings.
type Config interface {
	// GetDialect returns database dialect (e.g., "sqlite", "postgres").
	GetDialect() string

	// SetDialect sets the database dialect.
	SetDialect(dialect string) error

	// GetDSN returns the database connection string (Data Source Name).
	GetDSN() string

	// SetDSN sets the database connection string (Data Source Name).
	SetDSN(dsn string) error

	// IsDebug returns whether debug mode is enabled.
	IsDebug() bool

	// SetDebug enables or disables debug mode.
	SetDebug(debug bool) error
}
