package chatrole

import base "github.com/kylerqws/chatbot/pkg/openai/enumset/chatrole"

// Codes defines available chat role codes.
type Codes struct {
	System    string `json:"system"`
	User      string `json:"user"`
	Assistant string `json:"assistant"`
	Tool      string `json:"tool"`
}

// Manager provides access to available chat role values.
type Manager struct {
	List  map[string]*base.ChatRole
	Codes *Codes
}

// NewManager creates a new manager for available chat role values.
func NewManager() *Manager {
	return &Manager{
		List: base.AllChatRoles,
		Codes: &Codes{
			System:    base.SystemCode,
			User:      base.UserCode,
			Assistant: base.AssistantCode,
			Tool:      base.ToolCode,
		},
	}
}

// Resolve returns the chat role associated with the given code.
func (*Manager) Resolve(code string) (*base.ChatRole, error) {
	return base.Resolve(code)
}

// JoinCodes joins all known chat role codes using the given separator.
func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
