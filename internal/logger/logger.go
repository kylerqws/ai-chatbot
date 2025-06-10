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
func (m *manager) DB() ctrpkg.Logger {
	m.dbOnce.Do(func() {
		m.db = m.newLogger(ctrwrt.TypeDB)
	})
	return m.db
}

// Out returns the logger for standard output.
func (m *manager) Out() ctrpkg.Logger {
	m.stdoutOnce.Do(func() {
		m.stdout = m.newLogger(ctrwrt.TypeStdout)
	})
	return m.stdout
}

// Err returns the logger for standard error.
func (m *manager) Err() ctrpkg.Logger {
	m.stderrOnce.Do(func() {
		m.stderr = m.newLogger(ctrwrt.TypeStderr)
	})
	return m.stderr
}

// newLogger creates a logger for the specified writer type.
func (m *manager) newLogger(wt string) ctrpkg.Logger {
	instance, err := logger.NewWithWriter(m.ctx, wt)
	if err != nil {
		log.Fatal(fmt.Errorf("init logger '%s': %w", wt, err))
	}
	return instance
}
