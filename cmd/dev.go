package cmd

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd"
)

func DevCommand(app *intapp.App) *cobra.Command {
	return intcli.NewDevAdapter(app).Configure()
}
