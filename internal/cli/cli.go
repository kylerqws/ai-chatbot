package cli

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter"
)

// RootCommand creates the root command.
func RootCommand(app *intapp.App) *cobra.Command {
	// Disable automatic sorting of subcommands to preserve custom order.
	cobra.EnableCommandSorting = false

	return intcli.NewRootAdapter(app).Configure()
}

// Execute runs the CLI entry point by executing the root command.
func Execute(app *intapp.App) error {
	defer app.ContextCancel()()
	return RootCommand(app).Execute()
}
