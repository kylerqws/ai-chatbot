package service

import (
	"context"
	"encoding/json"

	"github.com/kylerqws/chatbot/pkg/openai/utils/value"
)

// ChatMessage represents a single message in a chat conversation.
type ChatMessage struct {
	Role       string      `json:"role"`
	Content    *string     `json:"content,omitempty"`
	Name       *string     `json:"name,omitempty"`
	ToolCallID *string     `json:"tool_call_id,omitempty"`
	ToolCalls  []*ToolCall `json:"tool_calls,omitempty"`
}

// ToolCall represents a tool invocation triggered by the model.
type ToolCall struct {
	ID       string        `json:"id"`
	Type     string        `json:"type"`
	Function *FunctionCall `json:"function"`
}

// Tool specifies a function available to the model for invocation.
type Tool struct {
	Type     string        `json:"type"`
	Function *FunctionSpec `json:"function"`
}

// FunctionCall represents a call to a function with given arguments.
type FunctionCall struct {
	Name      string           `json:"name"`
	Arguments *json.RawMessage `json:"arguments,omitempty"`
}

// FunctionSpec defines available function metadata and parameter schema.
type FunctionSpec struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
}

// ChatCompletionRequest is the request for generating chat completions.
type ChatCompletionRequest struct {
	Model            string            `json:"model"`
	User             *string           `json:"user,omitempty"`
	Stream           *bool             `json:"stream,omitempty"`
	MaxTokens        *int              `json:"max_tokens,omitempty"`
	Temperature      *float64          `json:"temperature,omitempty"`
	TopP             *float64          `json:"top_p,omitempty"`
	PresencePenalty  *float64          `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64          `json:"frequency_penalty,omitempty"`
	Stop             []string          `json:"stop,omitempty"`
	ToolChoice       *value.ToolChoice `json:"tool_choice,omitempty"`
	Messages         []*ChatMessage    `json:"messages"`
	Tools            []*Tool           `json:"tools,omitempty"`
}

// ChatCompletionChoice represents a single response choice in a chat completion response.
type ChatCompletionChoice struct {
	Index        int          `json:"index"`
	FinishReason string       `json:"finish_reason"`
	Message      *ChatMessage `json:"message"`
	ToolCalls    []*ToolCall  `json:"tool_calls,omitempty"`
}

// ChatCompletionResponse is the response returned from the OpenAI chat API.
type ChatCompletionResponse struct {
	ID      string                  `json:"id"`
	Object  string                  `json:"object"`
	Model   string                  `json:"model"`
	Created int64                   `json:"created"`
	Usage   *Usage                  `json:"usage,omitempty"`
	Choices []*ChatCompletionChoice `json:"choices"`
}

// Usage contains token usage statistics for a chat completion.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatService defines the operation for creating chat completions.
type ChatService interface {
	// ChatCompletion sends a chat completion request to OpenAI and returns the result.
	ChatCompletion(ctx context.Context, req *ChatCompletionRequest) (*ChatCompletionResponse, error)
}
