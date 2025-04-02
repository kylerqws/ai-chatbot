package config

import (
	"context"
	"fmt"

	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/config/source"
)

const (
	SourceTypeKey     = "configSourceType"
	DefaultSourceType = "env"
)

func New(ctx context.Context) (ctrcfg.Config, error) {
	st, ok := ctx.Value(SourceTypeKey).(string)
	if !ok || st == "" {
		st = DefaultSourceType
	}

	switch st {
	case "env":
		return source.NewEnvConfig(ctx)
	default:
		return nil, fmt.Errorf("unsupported OpenAI configuration source: %q", st)
	}
}
