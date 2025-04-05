package config

import (
	"context"
	"fmt"

	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/config/source"
)

const (
	SourceTypeKey     = "sourceType"
	DefaultSourceType = "env"
)

func New(ctx context.Context) (ctrcfg.Config, error) {
	st, ok := ctx.Value(SourceTypeKey).(string)
	if !ok || st == "" {
		st = DefaultSourceType
	}

	switch st {
	case "env":
		cfg, err := source.NewEnvConfig(ctx)
		if err != nil {
			return nil, fmt.Errorf("config: failed to load env source: %w", err)
		}
		return cfg, nil
	default:
		return nil, fmt.Errorf("config: unsupported source type: %q", st)
	}
}
