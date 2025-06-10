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

// manager provides access to internal loggers.
type manager struct {
	ctx context.Context

	db     ctrpkg.Logger
	dbOnce sync.Once

	stdout     ctrpkg.Logger
	stdoutOnce sync.Once

	stderr     ctrpkg.Logger
	stderrOnce sync.Once
}

// New returns a new logger manager.
func New(ctx context.Context) ctrint.Logger {
	return &manager{ctx: ctx}
}

// DB returns the logger for database output.
func (l *manager) DB() ctrpkg.Logger {
	l.dbOnce.Do(func() {
		l.db = l.initLogger(ctrwrt.TypeDB)
	})
	return l.db
}

// Out returns the logger for standard output.
func (l *manager) Out() ctrpkg.Logger {
	l.stdoutOnce.Do(func() {
		l.stdout = l.initLogger(ctrwrt.TypeStdout)
	})
	return l.stdout
}

// Err returns the logger for standard error.
func (l *manager) Err() ctrpkg.Logger {
	l.stderrOnce.Do(func() {
		l.stderr = l.initLogger(ctrwrt.TypeStderr)
	})
	return l.stderr
}

// initLogger initializes a logger for the specified writer type.
func (l *manager) initLogger(wt string) ctrpkg.Logger {
	instance, err := logger.NewWithWriter(l.ctx, wt)
	if err != nil {
		log.Fatal(fmt.Errorf("init logger '%s': %w", wt, err))
	}
	return instance
}
