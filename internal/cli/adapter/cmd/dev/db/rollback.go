package db

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"
	intmig "github.com/kylerqws/chatbot/internal/db/migrator"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// RollbackAdapter provides the implementation for the rollback command adapter.
type RollbackAdapter struct {
	*helper.CommandAdapter
}

// NewRollbackAdapter creates a new rollback command adapter.
func NewRollbackAdapter(app *intapp.App) ctr.CommandAdapter {
	adp := &RollbackAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapter = helper.NewCommandAdapter(app, cmd)
	return adp
}

// Configure applies configuration for the command.
func (a *RollbackAdapter) Configure() *cobra.Command {
	a.SetUse("rollback")
	a.SetShort("Rollback the most recent batch of database migrations")

	a.SetFuncRunE(a.Rollback)
	return a.MainConfigure()
}

// Rollback executes the database rollback process.
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
