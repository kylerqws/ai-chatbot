package chat

import (
	"context"

	ctrint "github.com/kylerqws/chatbot/internal/openai/contract/service"
	ctrpkg "github.com/kylerqws/chatbot/pkg/openai/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// service provides operations for generating chat completions.
type service struct {
	ctx context.Context
	sdk ctrpkg.OpenAI

	svc ctrsvc.ChatService
}

// NewService creates a new chat service for generating chat completions.
func NewService(ctx context.Context, sdk ctrpkg.OpenAI) ctrint.ChatService {
	return &service{ctx: ctx, sdk: sdk, svc: sdk.ChatService()}
}

// NewChatCompletionRequest creates a new chat completion request.
func (s *service) NewChatCompletionRequest() *ctrsvc.ChatCompletionRequest {
	return &ctrsvc.ChatCompletionRequest{}
}

// NewChatCompletionResponse creates a new chat completion response.
func (s *service) NewChatCompletionResponse() *ctrsvc.ChatCompletionResponse {
	return &ctrsvc.ChatCompletionResponse{}
}

// ChatCompletion sends a chat completion request and returns the response.
func (s *service) ChatCompletion(ctx context.Context, req *ctrsvc.ChatCompletionRequest) (*ctrsvc.ChatCompletionResponse, error) {
	return s.sdk.ChatService().ChatCompletion(ctx, req)
}
