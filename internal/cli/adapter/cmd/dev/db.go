package dev

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	"github.com/kylerqws/chatbot/cmd/dev/db"
	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type DBAdapter struct {
	app      *app.App
	command  *cobra.Command
	children []*cobra.Command
}

func NewDBAdapter(app *app.App) ctr.ParentAdapter {
	return &DBAdapter{app: app}
}

func (a *DBAdapter) Configure() *cobra.Command {
	a.command = &cobra.Command{
		Use:   "db",
		Short: "Commands for managing the database",
	}

	a.children = []*cobra.Command{
		db.MigrateCommand(a.app),
		db.RollbackCommand(a.app),
	}

	return setup.ParentConfigure(a)
}

func (a *DBAdapter) App() *app.App {
	return a.app
}

func (a *DBAdapter) Command() *cobra.Command {
	return a.command
}

func (a *DBAdapter) Children() []*cobra.Command {
	return a.children
}
