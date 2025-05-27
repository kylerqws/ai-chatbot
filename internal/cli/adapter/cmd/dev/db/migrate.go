package db

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"
	intmig "github.com/kylerqws/chatbot/internal/db/migrator"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type MigrateAdapter struct {
	*helper.CommandAdapter
}

func NewMigrateAdapter(app *intapp.App) ctr.CommandAdapter {
	adp := &MigrateAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapter = helper.NewCommandAdapter(app, cmd)
	return adp
}

func (a *MigrateAdapter) Configure() *cobra.Command {
	a.SetUse("migrate")
	a.SetShort("Run schema migrations for the application")
	a.SetFuncRunE(a.Migrate)

	return a.MainConfigure()
}

func (a *MigrateAdapter) Migrate(_ *cobra.Command, _ []string) error {
	app := a.App()

	err := intmig.Migrate(app.Context(), app.DB())
	if err != nil {
		a.AddError(err)
	}

	if !a.ExistErrors() {
		return a.PrintMessage("Database migrations applied successfully.")
	}
	return a.ErrorIfExist("failed to apply database migrations")
}
