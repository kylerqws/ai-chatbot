package openai

import (
	"context"

	"github.com/kylerqws/chatbot/pkg/openai"
	ctr "github.com/kylerqws/chatbot/pkg/openai/contract"
)

func New(ctx context.Context) (ctr.OpenAI, error) {
	return openai.New(ctx)
}
