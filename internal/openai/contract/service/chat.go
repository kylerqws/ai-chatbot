package service

import (
	"context"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// ChatService defines operations for generating chat completions via OpenAI.
type ChatService interface {
	// NewChatCompletionRequest creates a new chat completion request.
	NewChatCompletionRequest() *ctrsvc.ChatCompletionRequest

	// NewChatCompletionResponse creates a new chat completion response.
	NewChatCompletionResponse() *ctrsvc.ChatCompletionResponse

	// ChatCompletion sends a chat completion request and returns the response.
	ChatCompletion(ctx context.Context, req *ctrsvc.ChatCompletionRequest) (*ctrsvc.ChatCompletionResponse, error)
}
