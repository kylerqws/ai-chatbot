package config

import (
	"context"
	"fmt"

	ctrcf "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/config/source"
)

const (
	SourceTypeKey     = "configSourceType"
	DefaultSourceType = "env"
)

func New(ctx context.Context) (ctrcf.Config, error) {
	st, ok := ctx.Value(SourceTypeKey).(string)
	if !ok || st == "" {
		st = DefaultSourceType
	}

	switch st {
	case "env":
		return source.NewEnvConfig(ctx)
	default:
		return nil, fmt.Errorf("unsupported OpenAI API configuration source: %q", st)
	}
}
