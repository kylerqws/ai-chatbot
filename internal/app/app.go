package app

import (
	"context"

	intdb "github.com/kylerqws/chatbot/internal/db"
	intlog "github.com/kylerqws/chatbot/internal/logger"
	intopenai "github.com/kylerqws/chatbot/internal/openai"

	ctrdb "github.com/kylerqws/chatbot/pkg/db/contract"
	ctrlog "github.com/kylerqws/chatbot/pkg/logger/contract"
	ctropenai "github.com/kylerqws/chatbot/pkg/openai/contract"
)

type App struct {
	db     ctrdb.DB
	logger ctrlog.Logger
	openai ctropenai.OpenAI

	context   context.Context
	ctxCancel context.CancelFunc
}

const (
	Name    = "AI ChatBot"
	Version = "1.0.0-dev"
)

const (
	ModeUtility = "utility"
	ModeService = "service"
	DefaultMode = ModeUtility
)

func New(ctx context.Context, ctxCancel context.CancelFunc) (*App, error) {
	db, err := intdb.New(ctx)
	if err != nil {
		return nil, err
	}

	logger, err := intlog.New(ctx)
	if err != nil {
		return nil, err
	}

	openai, err := intopenai.New(ctx)
	if err != nil {
		return nil, err
	}

	return &App{
		db:     db,
		logger: logger,
		openai: openai,

		context:   ctx,
		ctxCancel: ctxCancel,
	}, nil
}

func (*App) Name() string {
	return Name
}

func (*App) Version() string {
	return Version
}

func (app *App) DB() ctrdb.DB {
	return app.db
}

func (app *App) Logger() ctrlog.Logger {
	return app.logger
}

func (app *App) OpenAI() ctropenai.OpenAI {
	return app.openai
}

func (app *App) Context() context.Context {
	return app.context
}

func (app *App) ContextCancel() context.CancelFunc {
	return app.ctxCancel
}
