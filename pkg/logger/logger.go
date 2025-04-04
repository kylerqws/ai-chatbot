package logger

import (
	"context"

	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/logger"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/writer"

	ctrlog "github.com/kylerqws/chatbot/pkg/logger/contract/logger"
)

func NewLogger(ctx context.Context) (ctrlog.Logger, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, err
	}

	prv, err := writer.New(cfg)
	if err != nil {
		return nil, err
	}

	return logger.NewZeroLogger(cfg, prv.Writer()), nil
}
