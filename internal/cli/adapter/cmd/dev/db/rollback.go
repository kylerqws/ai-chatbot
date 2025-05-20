package db

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	inthlp "github.com/kylerqws/chatbot/internal/cli/helper"
	intmig "github.com/kylerqws/chatbot/internal/db/migrator"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type RollbackAdapter struct {
	*inthlp.CommandAdapterHelper
	*inthlp.PrintAdapterHelper
}

func NewRollbackAdapter(app *intapp.App) ctr.CommandAdapter {
	adp := &RollbackAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapterHelper = inthlp.NewCommandAdapterHelper(adp, app, cmd)
	adp.PrintAdapterHelper = inthlp.NewPrintAdapterHelper(cmd)

	return adp
}

func (a *RollbackAdapter) Configure() *cobra.Command {
	a.SetUse("rollback")
	a.SetShort("Revert the most recent set of migrations")
	a.SetFuncRunE(a.FuncRunE)

	return a.MainConfigure()
}

func (a *RollbackAdapter) FuncRunE(_ *cobra.Command, _ []string) error {
	app := a.App()

	err := intmig.Rollback(app.Context(), app.DB())
	if err != nil {
		a.AddError(err)
	}

	if !a.ExistErrors() {
		return a.PrintMessage("Database rollback completed successfully.")
	}
	if !a.ShowErrors() {
		return a.PrintMessage("Failed to complete database rollback.")
	}

	return a.PrintErrors()
}
