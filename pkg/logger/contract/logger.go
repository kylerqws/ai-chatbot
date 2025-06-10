package contract

import "context"

// Logger provides structured logging with context support.
type Logger interface {
	// Info logs an informational message.
	Info(args ...any)
	// InfoWithContext logs an informational message with context.
	InfoWithContext(ctx context.Context, args ...any)

	// Error logs an error message.
	Error(args ...any)
	// ErrorWithContext logs an error message with context.
	ErrorWithContext(ctx context.Context, args ...any)

	// Debug logs a debug-level message.
	Debug(args ...any)
	// DebugWithContext logs a debug-level message with context.
	DebugWithContext(ctx context.Context, args ...any)
}
