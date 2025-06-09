package model

import (
	"fmt"
	"sort"
	"strings"
)

// Model defines a machine learning model supported by the OpenAI API.
type Model struct {
	Code        string // Unique identifier for the model.
	Description string // Human-readable explanation of the model's purpose.
}

// Model code constants.
const (
	GPT35TurboCode = "gpt-3.5-turbo"
	GPT4Code       = "gpt-4"
)

// Predefined Model instances.
var (
	GPT35Turbo = &Model{
		Code:        GPT35TurboCode,
		Description: "Fast and cost-effective model for general tasks and fine-tuning.",
	}
	GPT4 = &Model{
		Code:        GPT4Code,
		Description: "High-performance model for complex tasks (not fine-tunable).",
	}
)

// AllModels maps all known model codes to their Model instances.
var AllModels = map[string]*Model{
	GPT35TurboCode: GPT35Turbo,
	GPT4Code:       GPT4,
}

// Resolve returns the Model associated with the given code.
// If the code is empty, GPT-3.5 Turbo is returned by default.
// Returns an error if the code is unrecognized.
func Resolve(code string) (*Model, error) {
	if code == "" {
		return GPT35Turbo, nil
	}
	if mdl, ok := AllModels[code]; ok {
		return mdl, nil
	}
	return nil, fmt.Errorf("model code is unknown")
}

// JoinCodes returns a sorted, delimited string of all known model codes.
func JoinCodes(sep string) string {
	codes := make([]string, 0, len(AllModels))
	for code := range AllModels {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return strings.Join(codes, sep)
}
