package db

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/dev/db"
)

// RollbackCommand creates the rollback command.
func RollbackCommand(app *intapp.App) *cobra.Command {
	return intcli.NewRollbackAdapter(app).Configure()
}
