package app

import (
	"context"

	"github.com/kylerqws/chatbot/internal/db"
	"github.com/kylerqws/chatbot/internal/logger"
	"github.com/kylerqws/chatbot/internal/openai"

	ctrdb "github.com/kylerqws/chatbot/internal/db/contract"
	ctrlog "github.com/kylerqws/chatbot/internal/logger/contract"
	ctropenai "github.com/kylerqws/chatbot/internal/openai/contract"
)

// App is the main application structure.
type App struct {
	db     ctrdb.DB
	logger ctrlog.Logger
	openai ctropenai.OpenAI

	context   context.Context
	ctxCancel context.CancelFunc
}

// Application metadata.
const (
	Name    = "AI ChatBot"
	Version = "1.0.0-dev"
)

// Application modes.
const (
	ModeUtility = "utility"
	ModeService = "service"
	DefaultMode = ModeUtility
)

// New creates a new App instance with its dependencies.
func New(ctx context.Context, ctxCancel context.CancelFunc) *App {
	return &App{
		db:     db.New(ctx),
		logger: logger.New(ctx),
		openai: openai.New(ctx),

		context:   ctx,
		ctxCancel: ctxCancel,
	}
}

// Name returns the application name.
func (*App) Name() string {
	return Name
}

// Version returns the application version.
func (*App) Version() string {
	return Version
}

// DB returns the database interface.
func (app *App) DB() ctrdb.DB {
	return app.db
}

// Logger returns the logger interface.
func (app *App) Logger() ctrlog.Logger {
	return app.logger
}

// OpenAI returns the OpenAI interface.
func (app *App) OpenAI() ctropenai.OpenAI {
	return app.openai
}

// Context returns the application context.
func (app *App) Context() context.Context {
	return app.context
}

// ContextCancel returns the application's context cancellation function.
func (app *App) ContextCancel() context.CancelFunc {
	return app.ctxCancel
}
