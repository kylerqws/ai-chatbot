package value

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ToolChoice is a wrapper for encoding/decoding the "tool_choice" field,
// which accepts either a string ("auto", "none") or a structured object
// specifying a function to call.
//
// Examples:
//
//	"auto"
//	"none"
//	{ "type": "function", "function": { "name": "my_func" } }
type ToolChoice struct {
	raw *json.RawMessage
}

// NewToolChoiceAuto returns a ToolChoice representing the string "auto".
func NewToolChoiceAuto() *ToolChoice {
	data, _ := json.Marshal("auto")
	raw := json.RawMessage(data)
	return &ToolChoice{raw: &raw}
}

// NewToolChoiceNone returns a ToolChoice representing the string "none".
func NewToolChoiceNone() *ToolChoice {
	data, _ := json.Marshal("none")
	raw := json.RawMessage(data)
	return &ToolChoice{raw: &raw}
}

// NewToolChoiceFunction returns a ToolChoice for a specific function call.
// The name parameter must be a non-empty function name.
func NewToolChoiceFunction(name string) (*ToolChoice, error) {
	if name == "" {
		return nil, errors.New("function name cannot be empty")
	}

	payload := map[string]any{
		"type": "function",
		"function": map[string]string{
			"name": name,
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal tool_choice function object: %w", err)
	}

	raw := json.RawMessage(data)
	return &ToolChoice{raw: &raw}, nil
}

// MarshalJSON implements the json.Marshaler interface.
func (tc *ToolChoice) MarshalJSON() ([]byte, error) {
	if tc.raw == nil {
		return []byte("null"), nil
	}
	return *tc.raw, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (tc *ToolChoice) UnmarshalJSON(data []byte) error {
	raw := json.RawMessage(data)
	tc.raw = &raw
	return nil
}

// Raw returns the raw JSON-encoded value of tool_choice, which may be a string or object.
func (tc *ToolChoice) Raw() *json.RawMessage {
	return tc.raw
}

// String returns the JSON string representation of the tool_choice value.
func (tc *ToolChoice) String() string {
	if tc.raw == nil {
		return ""
	}
	return string(*tc.raw)
}
