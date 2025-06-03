package model

import base "github.com/kylerqws/chatbot/pkg/openai/domain/model"

type Codes struct {
	GPT35 string `json:"gpt_35"`
	GPT4  string `json:"gpt_4"`
}

type Manager struct {
	List  map[string]*base.Model
	Codes *Codes
}

func NewManager() *Manager {
	return &Manager{
		List: base.AllModels,
		Codes: &Codes{
			GPT35: base.GPT35Code,
			GPT4:  base.GPT4Code,
		},
	}
}

func (*Manager) Resolve(code string) (*base.Model, error) {
	return base.Resolve(code)
}

func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}

func (*Manager) Default() *base.Model {
	return base.GPT35
}
