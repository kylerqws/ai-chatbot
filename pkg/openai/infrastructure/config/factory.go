package config

import (
	"context"
	"fmt"

	ctr "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/config/source"
)

// New returns a Config implementation based on the source type defined in the context.
// If no source type is provided, the default from the contract is used.
func New(ctx context.Context) (ctr.Config, error) {
	st, ok := ctx.Value(ctr.SourceTypeKey).(ctr.SourceType)
	if !ok || st == "" {
		st = ctr.DefaultSourceType
	}

	switch st {
	case ctr.EnvSourceType:
		cfg, err := source.NewEnvConfig(ctx)
		if err != nil {
			return nil, fmt.Errorf("load env config: %w", err)
		}
		return cfg, nil

	default:
		return nil, fmt.Errorf("unsupported config source: '%s'", st)
	}
}
