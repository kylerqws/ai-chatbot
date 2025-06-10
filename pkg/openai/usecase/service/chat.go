package service

import (
	"context"
	"encoding/json"
	"fmt"

	ctrcl "github.com/kylerqws/chatbot/pkg/openai/contract/client"
	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// chatService implements ChatService using the OpenAI API client.
type chatService struct {
	config ctrcfg.Config
	client ctrcl.Client
}

// NewChatService creates a new instance of ChatService.
func NewChatService(cl ctrcl.Client, cfg ctrcfg.Config) ctrsvc.ChatService {
	return &chatService{config: cfg, client: cl}
}

// ChatCompletion sends a chat completion request to OpenAI and returns the generated response.
func (s *chatService) ChatCompletion(ctx context.Context, req *ctrsvc.ChatCompletionRequest) (*ctrsvc.ChatCompletionResponse, error) {
	result := &ctrsvc.ChatCompletionResponse{}

	resp, err := s.client.RequestJSON(ctx, "POST", "/chat/completions", req)
	if err != nil {
		return nil, fmt.Errorf("send chat completion request: %w", err)
	}

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, fmt.Errorf("unmarshal chat completion response: %w", err)
	}

	return result, nil
}
