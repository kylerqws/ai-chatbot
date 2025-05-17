package logger

import (
	"context"
	"fmt"

	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/logger"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/writer"

	ctr "github.com/kylerqws/chatbot/pkg/logger/contract"
)

func New(ctx context.Context) (ctr.Logger, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	prv, err := writer.NewProvider(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create writer provider: %w", err)
	}

	return logger.NewZeroLogger(cfg, prv.Writer()), nil
}
