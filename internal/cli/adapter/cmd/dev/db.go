package dev

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/dev/db"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// DBAdapter provides the implementation for the database tools CLI adapter.
type DBAdapter struct {
	*helper.ParentAdapter
}

// NewDBAdapter creates a new DBAdapter adapter.
func NewDBAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &DBAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapter = helper.NewParentAdapter(app, cmd)
	return adp
}

// Configure applies configuration for the command.
func (a *DBAdapter) Configure() *cobra.Command {
	app := a.App()

	a.SetUse("db")
	a.SetShort("Manage database schema and application data")

	a.AddChildren(
		action.MigrateCommand(app),
		action.RollbackCommand(app),
	)

	return a.MainConfigure()
}
