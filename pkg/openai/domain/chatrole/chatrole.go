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
	if cr, ok := AllChatRoles[code]; ok {
		return cr, nil
	}
	return nil, fmt.Errorf("unknown chat role code: '%s'", code)
}

// JoinCodes returns all chat role codes joined by separator.
func JoinCodes(sep string) string {
	c := make([]string, 0, len(AllChatRoles))
	for code := range AllChatRoles {
		c = append(c, code)
	}
	sort.Strings(c)
	return strings.Join(c, sep)
}
