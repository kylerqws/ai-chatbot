package owner

import base "github.com/kylerqws/chatbot/pkg/openai/enumset/owner"

// Codes defines available owner codes.
type Codes struct {
	OpenAI       string `json:"openai"`
	System       string `json:"system"`
	Organization string `json:"organization"`
	User         string `json:"user"`
}

// Manager provides access to available owner values.
type Manager struct {
	List  map[string]*base.Owner
	Codes *Codes
}

// NewManager creates a new manager for available owner values.
func NewManager() *Manager {
	return &Manager{
		List: base.AllOwners,
		Codes: &Codes{
			OpenAI:       base.OpenAICode,
			System:       base.SystemCode,
			Organization: base.OrganizationCode,
			User:         base.UserCode,
		},
	}
}

// Resolve returns the owner associated with the given code.
func (*Manager) Resolve(code string) (*base.Owner, error) {
	return base.Resolve(code)
}

// JoinCodes joins all known owner codes using the given separator.
func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}

// Default returns the default owner.
func (*Manager) Default() *base.Owner {
	return base.Default()
}
