package purpose

import base "github.com/kylerqws/chatbot/pkg/openai/domain/purpose"

type Manager struct {
	List map[string]*base.Purpose
}

func NewManager() *Manager {
	return &Manager{List: base.AllPurposes}
}

func (*Manager) Resolve(code string) (*base.Purpose, error) {
	return base.Resolve(code)
}

func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}

func (*Manager) Default() *base.Purpose {
	return base.FineTune
}
