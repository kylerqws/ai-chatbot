package logger

import (
	"context"

	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/logger"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/registry"

	ctrlog "github.com/kylerqws/chatbot/pkg/logger/contract/logger"
	ctrreg "github.com/kylerqws/chatbot/pkg/logger/contract/registry"
)

func NewRegistry(ctx context.Context) (ctrreg.LoggerRegistry, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, err
	}

	reg := registry.New()
	reg.Register("default", logger.NewZeroLogger(cfg))

	return reg, nil
}

func NewDefaultLogger(ctx context.Context) (ctrlog.Logger, error) {
	reg, err := NewRegistry(ctx)
	if err != nil {
		return nil, err
	}

	return reg.Logger("default")
}
