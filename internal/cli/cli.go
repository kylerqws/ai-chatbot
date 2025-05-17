package cli

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	cli "github.com/kylerqws/chatbot/internal/cli/adapter"
)

func RootCommand(app *app.App) *cobra.Command {
	cobra.EnableCommandSorting = false
	return cli.NewRootAdapter(app).Configure()
}

func Execute(app *app.App) error {
	return RootCommand(app).Execute()
}
