package purpose

import base "github.com/kylerqws/chatbot/pkg/openai/enumset/purpose"

// Codes defines available purpose codes.
type Codes struct {
	FineTune   string `json:"fine_tune"`
	Assistants string `json:"assistants"`
	Batch      string `json:"batch"`
	UserData   string `json:"user_data"`
	Vision     string `json:"vision"`
	Evals      string `json:"evals"`
}

// Manager provides access to available purpose values.
type Manager struct {
	List  map[string]*base.Purpose
	Codes *Codes
}

// NewManager returns a new purpose manager.
func NewManager() *Manager {
	return &Manager{
		List: base.AllPurposes,
		Codes: &Codes{
			FineTune:   base.FineTuneCode,
			Assistants: base.AssistantsCode,
			Batch:      base.BatchCode,
			UserData:   base.UserDataCode,
			Vision:     base.VisionCode,
			Evals:      base.EvalsCode,
		},
	}
}

// Resolve returns the purpose associated with the given code.
func (*Manager) Resolve(code string) (*base.Purpose, error) {
	return base.Resolve(code)
}

// JoinCodes joins all known purpose codes using the given separator.
func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}

// Default returns the default purpose.
func (*Manager) Default() *base.Purpose {
	return base.Default()
}
