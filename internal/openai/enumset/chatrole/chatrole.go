package chatrole

import base "github.com/kylerqws/chatbot/pkg/openai/domain/chatrole"

type Manager struct {
	List map[string]*base.ChatRole
}

func NewManager() *Manager {
	return &Manager{List: base.AllChatRoles}
}

func (*Manager) Resolve(code string) (*base.ChatRole, error) {
	return base.Resolve(code)
}

func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
