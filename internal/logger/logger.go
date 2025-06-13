package logger

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/kylerqws/chatbot/pkg/logger"

	ctrint "github.com/kylerqws/chatbot/internal/logger/contract"
	ctrpkg "github.com/kylerqws/chatbot/pkg/logger/contract"
	ctrwrt "github.com/kylerqws/chatbot/pkg/logger/contract/writer"
)

// entrypoint aggregates internal loggers (database, stdout, stderr).
type entrypoint struct {
	ctx context.Context

	db     ctrpkg.Logger
	dbOnce sync.Once

	out     ctrpkg.Logger
	outOnce sync.Once

	err     ctrpkg.Logger
	errOnce sync.Once
}

// New creates a new logger entrypoint with multiple output loggers.
func New(ctx context.Context) ctrint.Logger {
	return &entrypoint{ctx: ctx}
}

// DB returns the logger for database output.
func (e *entrypoint) DB() ctrpkg.Logger {
	e.dbOnce.Do(func() {
		e.db = e.newLogger(ctrwrt.TypeDB)
	})
	return e.db
}

// Out returns the logger for standard output.
func (e *entrypoint) Out() ctrpkg.Logger {
	e.outOnce.Do(func() {
		e.out = e.newLogger(ctrwrt.TypeStdout)
	})
	return e.out
}

// Err returns the logger for standard error.
func (e *entrypoint) Err() ctrpkg.Logger {
	e.errOnce.Do(func() {
		e.err = e.newLogger(ctrwrt.TypeStderr)
	})
	return e.err
}

// newLogger creates a logger for the specified writer type.
func (e *entrypoint) newLogger(wt string) ctrpkg.Logger {
	instance, err := logger.NewWithWriter(e.ctx, wt)
	if err != nil {
		log.Fatal(fmt.Errorf("create logger with writer '%s': %w", wt, err))
	}
	return instance
}
