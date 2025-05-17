package file

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	cli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai/file"
)

func ListCommand(app *app.App) *cobra.Command {
	return cli.NewListAdapter(app).Configure()
}
