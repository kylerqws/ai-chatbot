package model

import "github.com/kylerqws/chatbot/pkg/openai/utils/enumset"

// Model defines a machine learning model supported by the OpenAI API.
type Model struct {
	Code        string // Unique identifier for the model.
	Description string // Human-readable explanation of the model.
}

// Model code constants.
const (
	GPT35TurboCode = "gpt‑3.5‑turbo"
	GPT4Code       = "gpt‑4"
)

// Predefined Model instances.
var (
	GPT35Turbo = &Model{Code: GPT35TurboCode, Description: "Fast and cost‑effective for general tasks."}
	GPT4       = &Model{Code: GPT4Code, Description: "High‑performance model for complex tasks."}
)

// AllModels lists all known Model instances.
var AllModels = map[string]*Model{
	GPT35TurboCode: GPT35Turbo,
	GPT4Code:       GPT4,
}

// Resolve looks up a Model by code, defaulting to GPT35Turbo.
func Resolve(code string) (*Model, error) {
	return enumset.ResolveWithDefault(code, AllModels, GPT35Turbo, "model")
}

// JoinCodes returns all model codes joined by separator.
func JoinCodes(sep string) string {
	return enumset.JoinCodes(AllModels, sep)
}
