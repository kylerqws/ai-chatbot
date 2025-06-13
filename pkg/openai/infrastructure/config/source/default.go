package source

const (
	// DefaultBaseURL is the default OpenAI API base URL.
	DefaultBaseURL = "https://api.openai.com/v1"

	// DefaultTimeout is the default HTTP client timeout in seconds.
	DefaultTimeout uint64 = 30

	// DefaultTLSHandshakeTimeout is the default TLS handshake timeout in seconds.
	DefaultTLSHandshakeTimeout uint64 = 10

	// DefaultResponseHeaderTimeout is the default response header timeout in seconds.
	DefaultResponseHeaderTimeout uint64 = 10

	// DefaultAPIKey is the default value for the API key.
	DefaultAPIKey = ""
)
