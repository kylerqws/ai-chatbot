package chatrole

import (
	"fmt"
	"sort"
	"strings"
)

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

// Resolve looks up a ChatRole by code, error if missing or unknown.
func Resolve(code string) (*ChatRole, error) {
	if code == "" {
		return nil, fmt.Errorf("chat role code is required")
	}
	if role, ok := AllChatRoles[code]; ok {
		return role, nil
	}
	return nil, fmt.Errorf("unknown chat role code: '%s'", code)
}

// JoinCodes returns all chat role codes joined by separator.
func JoinCodes(sep string) string {
	codes := make([]string, 0, len(AllChatRoles))
	for code := range AllChatRoles {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return strings.Join(codes, sep)
}
