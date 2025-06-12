package model

import base "github.com/kylerqws/chatbot/pkg/openai/enumset/model"

// Codes defines available model codes.
type Codes struct {
	GPT35Turbo string `json:"gpt_35_turbo"`
	GPT4       string `json:"gpt_4"`
}

// Manager provides access to available model values.
type Manager struct {
	List  map[string]*base.Model
	Codes *Codes
}

// NewManager returns a new model manager.
func NewManager() *Manager {
	return &Manager{
		List: base.AllModels,
		Codes: &Codes{
			GPT35Turbo: base.GPT35TurboCode,
			GPT4:       base.GPT4Code,
		},
	}
}

// Resolve returns the model associated with the given code.
func (*Manager) Resolve(code string) (*base.Model, error) {
	return base.Resolve(code)
}

// JoinCodes joins all known model codes using the given separator.
func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}

// Default returns the default model.
func (*Manager) Default() *base.Model {
	return base.Default()
}
