package config

import (
	"context"
	"fmt"

	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"
	"github.com/kylerqws/chatbot/pkg/db/infrastructure/config/source"
)

// New returns a Config implementation based on the source type defined in the context.
// If no source type is provided, the default from the contract is used.
func New(ctx context.Context) (ctrcfg.Config, error) {
	st, ok := ctx.Value(ctrcfg.SourceTypeKey).(ctrcfg.SourceType)
	if !ok || st == "" {
		st = ctrcfg.DefaultSourceType
	}

	switch st {
	case ctrcfg.EnvSourceType:
		cfg, err := source.NewEnvConfig(ctx)
		if err != nil {
			return nil, fmt.Errorf("load env config: %w", err)
		}
		return cfg, nil
	default:
		return nil, fmt.Errorf("unsupported config source: '%s'", st)
	}
}
