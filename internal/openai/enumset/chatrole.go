package enumset

import base "github.com/kylerqws/chatbot/pkg/openai/domain/chatrole"

type ChatRoleManager struct {
	List map[string]*base.ChatRole
}

func NewChatRoleManager() *ChatRoleManager {
	return &ChatRoleManager{List: base.AllChatRoles}
}

func (*ChatRoleManager) Resolve(code string) (*base.ChatRole, error) {
	return base.Resolve(code)
}

func (*ChatRoleManager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
