package logger

import (
	"context"
	"fmt"

	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/logger"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/writer"

	ctr "github.com/kylerqws/chatbot/pkg/logger/contract"
	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
)

// New returns a logger configured using values from the default config source.
func New(ctx context.Context) (ctr.Logger, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("load logger config: %w", err)
	}
	return newLogger(cfg)
}

// NewWithWriter returns a logger configured with the specified writer type.
func NewWithWriter(ctx context.Context, writerType string) (ctr.Logger, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("load logger config: %w", err)
	}

	if err := cfg.SetWriter(writerType); err != nil {
		return nil, fmt.Errorf("set writer: %w", err)
	}
	return newLogger(cfg)
}

// newLogger creates a logger using the provided configuration.
func newLogger(cfg ctrcfg.Config) (ctr.Logger, error) {
	prov, err := writer.NewProvider(cfg)
	if err != nil {
		return nil, fmt.Errorf("create writer provider: %w", err)
	}
	return logger.NewZeroLogger(cfg, prov.Writer()), nil
}
