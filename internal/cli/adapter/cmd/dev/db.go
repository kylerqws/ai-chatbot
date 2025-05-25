package dev

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/dev/db"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type DBAdapter struct {
	*helper.ParentAdapterHelper
}

func NewDBAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &DBAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper = helper.NewParentAdapterHelper(app, cmd)
	return adp
}

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
