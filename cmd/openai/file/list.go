package file

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai/file"
)

func ListCommand(app *intapp.App) *cobra.Command {
	return intcli.NewListAdapter(app).Configure()
}
