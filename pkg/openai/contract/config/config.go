package config

import "time"

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

// Config defines access to OpenAI API client settings.
type Config interface {
	// GetBaseURL returns the base URL of the OpenAI API.
	GetBaseURL() string

	// SetBaseURL sets the base URL for the API.
	SetBaseURL(url string) error

	// GetAPIKey returns the API key used for authentication.
	GetAPIKey() string

	// SetAPIKey sets the API key for authentication.
	SetAPIKey(key string) error

	// GetTimeout returns the HTTP timeout duration.
	GetTimeout() time.Duration

	// SetTimeout sets the timeout in seconds.
	SetTimeout(seconds uint64) error
}
