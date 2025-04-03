package logger

import "context"

type Logger interface {
	Info(...any)
	InfoWithContext(ctx context.Context, args ...any)

	Error(...any)
	ErrorWithContext(ctx context.Context, args ...any)

	Debug(...any)
	DebugWithContext(ctx context.Context, args ...any)
}
