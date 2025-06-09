package owner

import (
	"fmt"
	"sort"
	"strings"
)

// Owner represents a model ownership category used in OpenAI operations.
type Owner struct {
	Code        string // Unique identifier for the owner.
	Description string // Human-readable explanation of the owner type.
}

// Predefined owner codes.
const (
	OpenAICode       = "openai"
	SystemCode       = "system"
	OrganizationCode = "organization"
	UserCode         = "user"
)

// Predefined owner instances.
var (
	OpenAI = &Owner{
		Code:        OpenAICode,
		Description: "Owned by OpenAI and publicly available.",
	}
	System = &Owner{
		Code:        SystemCode,
		Description: "Used internally for system-level models.",
	}
	Organization = &Owner{
		Code:        OrganizationCode,
		Description: "Owned by the user's organization.",
	}
	User = &Owner{
		Code:        UserCode,
		Description: "Created by the authenticated user.",
	}
)

// AllOwners provides a lookup map from owner codes to their Owner definitions.
var AllOwners = map[string]*Owner{
	OpenAICode:       OpenAI,
	SystemCode:       System,
	OrganizationCode: Organization,
	UserCode:         User,
}

// Resolve returns the Owner associated with the given code.
// If the code is empty, it defaults to the User owner.
// Returns an error if the code is unrecognized.
func Resolve(code string) (*Owner, error) {
	if code == "" {
		return User, nil
	}
	if o, ok := AllOwners[code]; ok {
		return o, nil
	}
	return nil, fmt.Errorf("owner code is unknown")
}

// JoinCodes returns a string of all known owner codes joined by the specified separator.
func JoinCodes(sep string) string {
	codes := make([]string, 0, len(AllOwners))
	for code := range AllOwners {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return strings.Join(codes, sep)
}
