package db

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"
	intmig "github.com/kylerqws/chatbot/internal/db/migrator"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type RollbackAdapter struct {
	*helper.CommandAdapter
}

func NewRollbackAdapter(app *intapp.App) ctr.CommandAdapter {
	adp := &RollbackAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapter = helper.NewCommandAdapter(app, cmd)
	return adp
}

func (a *RollbackAdapter) Configure() *cobra.Command {
	a.SetUse("rollback")
	a.SetShort("Rollback the last set of migrations")
	a.SetFuncRunE(a.Rollback)

	return a.MainConfigure()
}

func (a *RollbackAdapter) Rollback(_ *cobra.Command, _ []string) error {
	app := a.App()

	err := intmig.Rollback(app.Context(), app.DB())
	if err != nil {
		a.AddError(err)
	}

	if !a.ExistErrors() {
		return a.PrintMessage("Database rollback completed successfully.")
	}
	return a.ErrorIfExist("Failed to complete database rollback.")
}
