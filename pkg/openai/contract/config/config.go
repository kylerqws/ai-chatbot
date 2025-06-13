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
	// GetBaseURL returns the OpenAI API base URL.
	GetBaseURL() string

	// SetBaseURL sets the OpenAI API base URL.
	SetBaseURL(url string) error

	// GetAPIKey returns the API key used for authentication.
	GetAPIKey() string

	// SetAPIKey sets the API key for authentication.
	SetAPIKey(key string) error

	// GetTimeout returns the HTTP client timeout duration.
	GetTimeout() time.Duration

	// SetTimeout sets the HTTP client timeout in seconds.
	SetTimeout(seconds uint64) error

	// GetTLSHandshakeTimeout returns the TLS handshake timeout duration.
	GetTLSHandshakeTimeout() time.Duration

	// SetTLSHandshakeTimeout sets the TLS handshake timeout in seconds.
	SetTLSHandshakeTimeout(seconds uint64) error

	// GetResponseHeaderTimeout returns the response header timeout duration.
	GetResponseHeaderTimeout() time.Duration

	// SetResponseHeaderTimeout sets the response header timeout in seconds.
	SetResponseHeaderTimeout(seconds uint64) error
}
