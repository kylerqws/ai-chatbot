package eventlevel

import base "github.com/kylerqws/chatbot/pkg/openai/enumset/eventlevel"

// Codes defines available event level codes.
type Codes struct {
	Info    string `json:"info"`
	Warning string `json:"warning"`
	Error   string `json:"error"`
}

// Manager provides access to available event level values.
type Manager struct {
	List  map[string]*base.EventLevel
	Codes *Codes
}

// NewManager returns a new event level manager.
func NewManager() *Manager {
	return &Manager{
		List: base.AllEventLevels,
		Codes: &Codes{
			Info:    base.InfoCode,
			Warning: base.WarningCode,
			Error:   base.ErrorCode,
		},
	}
}

// Resolve returns the event level associated with the given code.
func (*Manager) Resolve(code string) (*base.EventLevel, error) {
	return base.Resolve(code)
}

// JoinCodes joins all known event level codes using the given separator.
func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
