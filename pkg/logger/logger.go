package logger

import (
	"context"
	"fmt"

	ctr "github.com/kylerqws/chatbot/pkg/logger/contract"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/logger"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/writer"
)

func New(ctx context.Context) (ctr.Logger, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("logger: failed to load config: %w", err)
	}

	prv, err := writer.NewProvider(cfg)
	if err != nil {
		return nil, fmt.Errorf("logger: failed to create writer provider: %w", err)
	}

	return logger.NewZeroLogger(cfg, prv.Writer()), nil
}
