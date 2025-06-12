package chatrole

import "github.com/kylerqws/chatbot/pkg/openai/utils/enumset"

// ChatRole defines a role in chat conversations.
type ChatRole struct {
	Code        string // Unique identifier for the chat role.
	Description string // Human-readable explanation of the role.
}

// ChatRole code constants.
const (
	SystemCode    = "system"
	UserCode      = "user"
	AssistantCode = "assistant"
	ToolCode      = "tool"
)

// Predefined ChatRole instances.
var (
	System    = &ChatRole{Code: SystemCode, Description: "System instructions guiding behavior."}
	User      = &ChatRole{Code: UserCode, Description: "User input in conversation."}
	Assistant = &ChatRole{Code: AssistantCode, Description: "Assistant-generated response."}
	Tool      = &ChatRole{Code: ToolCode, Description: "Tool or function call output."}
)

// AllChatRoles lists all known ChatRole instances.
var AllChatRoles = map[string]*ChatRole{
	SystemCode:    System,
	UserCode:      User,
	AssistantCode: Assistant,
	ToolCode:      Tool,
}

// Resolve returns the ChatRole for the given code or error if missing or unknown.
func Resolve(code string) (*ChatRole, error) {
	return enumset.ResolveRequired(code, AllChatRoles, "chat role")
}

// JoinCodes returns all chat role codes joined by separator.
func JoinCodes(sep string) string {
	return enumset.JoinCodes(AllChatRoles, sep)
}
