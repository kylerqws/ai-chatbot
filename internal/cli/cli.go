package cli

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/adapter"
)

// RootCommand creates the root CLI command.
func RootCommand(app *app.App) *cobra.Command {
	// Disable automatic sorting of subcommands to preserve custom order.
	cobra.EnableCommandSorting = false

	return adapter.NewRootAdapter(app).Configure()
}

// Execute runs the root CLI command.
func Execute(app *app.App) error {
	defer app.ContextCancel()()
	return RootCommand(app).Execute()
}
