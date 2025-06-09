package owner

import (
	"fmt"
	"sort"
	"strings"
)

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

// Resolve looks up an Owner by code, defaulting to User.
func Resolve(code string) (*Owner, error) {
	if code == "" {
		return User, nil
	}
	if owner, ok := AllOwners[code]; ok {
		return owner, nil
	}
	return nil, fmt.Errorf("unknown owner code: '%s'", code)
}

// JoinCodes returns all owner codes joined by separator.
func JoinCodes(sep string) string {
	codes := make([]string, 0, len(AllOwners))
	for code := range AllOwners {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return strings.Join(codes, sep)
}
