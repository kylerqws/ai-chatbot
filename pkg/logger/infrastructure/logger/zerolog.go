package logger

import (
	"context"
	"fmt"
	"io"

	"github.com/rs/zerolog"

	ctr "github.com/kylerqws/chatbot/pkg/logger/contract"
	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
)

// zeroLogger is a zerolog-based implementation of the Logger interface.
type zeroLogger struct {
	config ctrcfg.Config
	logger *zerolog.Logger
}

// NewZeroLogger creates and returns a new zerolog-based logger instance.
func NewZeroLogger(cfg ctrcfg.Config, w io.Writer) ctr.Logger {
	zl := zerolog.New(w).With().Timestamp().Logger()
	return &zeroLogger{config: cfg, logger: &zl}
}

// Info logs an informational message.
func (l *zeroLogger) Info(args ...any) {
	l.logger.Info().Msg(l.format(args...))
}

// InfoWithContext logs an informational message with context.
func (l *zeroLogger) InfoWithContext(ctx context.Context, args ...any) {
	l.withContext(ctx).Info().Msg(l.format(args...))
}

// Error logs an error message.
func (l *zeroLogger) Error(args ...any) {
	l.logger.Error().Msg(l.format(args...))
}

// ErrorWithContext logs an error message with context.
func (l *zeroLogger) ErrorWithContext(ctx context.Context, args ...any) {
	l.withContext(ctx).Error().Msg(l.format(args...))
}

// Debug logs a debug-level message.
func (l *zeroLogger) Debug(args ...any) {
	if l.config.IsDebug() {
		l.logger.Debug().Msg(l.format(args...))
	}
}

// DebugWithContext logs a debug-level message with context.
func (l *zeroLogger) DebugWithContext(ctx context.Context, args ...any) {
	if l.config.IsDebug() {
		l.withContext(ctx).Debug().Msg(l.format(args...))
	}
}

// withContext returns a logger enriched with context, if available.
func (l *zeroLogger) withContext(ctx context.Context) *zerolog.Logger {
	if ctx == nil {
		return l.logger
	}
	if ctxLogger := zerolog.Ctx(ctx); ctxLogger != nil {
		logger := ctxLogger.With().Logger()
		return &logger
	}
	return l.logger
}

// format joins arguments into a single message string.
func (*zeroLogger) format(args ...any) string {
	if len(args) == 1 {
		if str, ok := args[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(args...)
}
