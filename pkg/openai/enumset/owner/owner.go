package owner

import "github.com/kylerqws/chatbot/pkg/openai/utils/enumset"

// Owner represents a model ownership category in OpenAI operations.
type Owner struct {
	Code        string // Unique identifier for the owner.
	Description string // Human-readable explanation of the owner type.
}

// Owner code constants.
const (
	OpenAICode       = "openai"
	SystemCode       = "system"
	OrganizationCode = "organization"
	UserCode         = "user"
)

// Predefined Owner instances.
var (
	OpenAI       = &Owner{Code: OpenAICode, Description: "Owned by OpenAI and publicly available."}
	System       = &Owner{Code: SystemCode, Description: "Used internally for system‑level models."}
	Organization = &Owner{Code: OrganizationCode, Description: "Owned by the user's organization."}
	User         = &Owner{Code: UserCode, Description: "Created by the authenticated user."}
)

// AllOwners lists all known Owner instances.
var AllOwners = map[string]*Owner{
	OpenAICode:       OpenAI,
	SystemCode:       System,
	OrganizationCode: Organization,
	UserCode:         User,
}

// Resolve returns the Owner for the given code or the default Owner if not found.
func Resolve(code string) (*Owner, error) {
	return enumset.ResolveWithDefault(code, AllOwners, Default(), "owner")
}

// JoinCodes returns all owner codes joined by separator.
func JoinCodes(sep string) string {
	return enumset.JoinCodes(AllOwners, sep)
}

// Default returns the default Owner.
func Default() *Owner {
	return User
}
