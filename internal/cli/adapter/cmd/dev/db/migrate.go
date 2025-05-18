package db

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	inthlp "github.com/kylerqws/chatbot/internal/cli/helper"
	intmig "github.com/kylerqws/chatbot/internal/db/migrator"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type MigrateAdapter struct {
	*inthlp.CommandAdapterHelper
	*inthlp.PrintAdapterHelper
}

func NewMigrateAdapter(app *intapp.App) ctr.CommandAdapter {
	adp := &MigrateAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapterHelper = inthlp.NewCommandAdapterHelper(adp, app, cmd)
	adp.PrintAdapterHelper = inthlp.NewPrintAdapterHelper(cmd)

	return adp
}

func (a *MigrateAdapter) Configure() *cobra.Command {
	a.SetUse("migrate")
	a.SetShort("Run database schema migrations for the application")
	a.SetFuncRunE(a.FuncRunE)

	return a.MainConfigure()
}

func (a *MigrateAdapter) FuncRunE(_ *cobra.Command, _ []string) error {
	app := a.App()

	err := intmig.Migrate(app.Context(), app.DB())
	if err != nil {
		a.AddError(err)
	}

	if !a.ExistErrors() {
		return a.PrintMessage("Database migrations have been applied successfully.")
	}
	if !a.ShowErrors() {
		return a.PrintMessage("Database migrations have not been applied.")
	}

	return a.PrintErrors()
}
