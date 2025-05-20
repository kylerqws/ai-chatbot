package db

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/dev/db"
)

func MigrateCommand(app *intapp.App) *cobra.Command {
	return intcli.NewMigrateAdapter(app).Configure()
}
