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
	svc ctrsvc.ChatService
}

// NewService creates a new chat service for generating chat completions.
func NewService(ctx context.Context, sdk ctrpkg.OpenAI) ctrint.ChatService {
	return &service{ctx: ctx, svc: sdk.ChatService()}
}

// NewChatCompletionRequest creates a new chat completion request.
func (*service) NewChatCompletionRequest() *ctrsvc.ChatCompletionRequest {
	return &ctrsvc.ChatCompletionRequest{}
}

// NewChatCompletionResponse creates a new chat completion response.
func (*service) NewChatCompletionResponse() *ctrsvc.ChatCompletionResponse {
	return &ctrsvc.ChatCompletionResponse{}
}

// ChatCompletion sends a chat completion request to OpenAI.
func (s *service) ChatCompletion(ctx context.Context, req *ctrsvc.ChatCompletionRequest) (*ctrsvc.ChatCompletionResponse, error) {
	return s.svc.ChatCompletion(ctx, req)
}
