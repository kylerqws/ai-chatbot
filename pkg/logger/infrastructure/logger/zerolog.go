package logger

import (
	"context"
	"fmt"
	"io"

	"github.com/rs/zerolog"

	ctr "github.com/kylerqws/chatbot/pkg/logger/contract"
	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
)

type zeroLogger struct {
	config ctrcfg.Config
	logger *zerolog.Logger
}

func NewZeroLogger(cfg ctrcfg.Config, w io.Writer) ctr.Logger {
	zl := zerolog.New(w).With().Timestamp().Logger()
	return &zeroLogger{config: cfg, logger: &zl}
}

func (l *zeroLogger) Info(args ...any) {
	l.logger.Info().Msg(l.format(args...))
}

func (l *zeroLogger) InfoWithContext(ctx context.Context, args ...any) {
	l.from(ctx).Info().Msg(l.format(args...))
}

func (l *zeroLogger) Error(args ...any) {
	l.logger.Error().Msg(l.format(args...))
}

func (l *zeroLogger) ErrorWithContext(ctx context.Context, args ...any) {
	l.from(ctx).Error().Msg(l.format(args...))
}

func (l *zeroLogger) Debug(args ...any) {
	if l.config.IsDebug() {
		l.logger.Debug().Msg(l.format(args...))
	}
}

func (l *zeroLogger) DebugWithContext(ctx context.Context, args ...any) {
	if l.config.IsDebug() {
		l.from(ctx).Debug().Msg(l.format(args...))
	}
}

func (l *zeroLogger) from(ctx context.Context) *zerolog.Logger {
	if ctx == nil {
		return l.logger
	}

	if ctxLogger := zerolog.Ctx(ctx); ctxLogger != nil {
		logger := ctxLogger.With().Logger()
		return &logger
	}

	return l.logger
}

func (l *zeroLogger) format(args ...any) string {
	if len(args) == 1 {
		if str, ok := args[0].(string); ok {
			return str
		}
	}

	return fmt.Sprint(args...)
}
