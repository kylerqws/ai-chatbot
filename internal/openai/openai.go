package openai

import (
	"context"

	"github.com/kylerqws/chatbot/pkg/openai"
	ctrapi "github.com/kylerqws/chatbot/pkg/openai/contract"
)

func New(ctx context.Context) (ctrapi.OpenAI, error) {
	return openai.New(ctx)
}
