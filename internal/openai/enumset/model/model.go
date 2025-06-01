package model

import base "github.com/kylerqws/chatbot/pkg/openai/domain/model"

type Manager struct {
	List map[string]*base.Model
}

func NewManager() *Manager {
	return &Manager{List: base.AllModels}
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
