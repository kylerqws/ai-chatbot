package logger

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"

	ctrc "github.com/kylerqws/chatbot/pkg/logger/contract/config"
	ctrl "github.com/kylerqws/chatbot/pkg/logger/contract/logger"
)

type zeroLogger struct {
	config ctrc.Config
	logger *zerolog.Logger
}

func NewZeroLogger(cfg ctrc.Config) ctrl.Logger {
	zl := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &zeroLogger{config: cfg, logger: &zl}
}

func (z *zeroLogger) Info(args ...any) {
	z.logger.Info().Msg(z.format(args...))
}

func (z *zeroLogger) InfoWithContext(ctx context.Context, args ...any) {
	z.from(ctx).Info().Msg(z.format(args...))
}

func (z *zeroLogger) Error(args ...any) {
	z.logger.Error().Msg(z.format(args...))
}

func (z *zeroLogger) ErrorWithContext(ctx context.Context, args ...any) {
	z.from(ctx).Error().Msg(z.format(args...))
}

func (z *zeroLogger) Debug(args ...any) {
	if z.config.IsDebug() {
		z.logger.Debug().Msg(z.format(args...))
	}
}

func (z *zeroLogger) DebugWithContext(ctx context.Context, args ...any) {
	if z.config.IsDebug() {
		z.from(ctx).Debug().Msg(z.format(args...))
	}
}

func (z *zeroLogger) from(ctx context.Context) *zerolog.Logger {
	if ctx == nil {
		return z.logger
	}

	if ctxLogger := zerolog.Ctx(ctx); ctxLogger != nil {
		logger := ctxLogger.With().Logger()
		return &logger
	}

	return z.logger
}

func (z *zeroLogger) format(args ...any) string {
	if len(args) == 1 {
		if str, ok := args[0].(string); ok {
			return str
		}
	}

	return fmt.Sprint(args...)
}
