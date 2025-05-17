package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"

	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type chatService struct {
	config ctrcfg.Config
	client *client.Client
}

func NewChatService(cl *client.Client, cfg ctrcfg.Config) ctrsvc.ChatService {
	return &chatService{config: cfg, client: cl}
}

func (s *chatService) ChatCompletion(
	ctx context.Context,
	req *ctrsvc.ChatCompletionRequest,
) (*ctrsvc.ChatCompletionResponse, error) {
	result := &ctrsvc.ChatCompletionResponse{}

	payload, err := json.Marshal(req)
	if err != nil {
		return result, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.RequestReader(ctx, "POST", "/chat/completions", bytes.NewReader(payload))
	if err != nil {
		return result, fmt.Errorf("failed to send request: %w", err)
	}

	err = json.Unmarshal(resp, result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return result, nil
}
