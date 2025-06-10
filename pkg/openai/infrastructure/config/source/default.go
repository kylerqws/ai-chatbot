package source

const (
	// DefaultBaseURL is the default OpenAI API base URL.
	DefaultBaseURL = "https://api.openai.com/v1"

	// DefaultTimeout is the default HTTP client timeout in seconds.
	DefaultTimeout uint64 = 30

	// DefaultAPIKey is the default value for the API key.
	// This value must be explicitly overridden by the selected config source.
	DefaultAPIKey = ""
)
