package file

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai/model"
)

// ListCommand creates a command for listing OpenAI models.
func ListCommand(app *intapp.App) *cobra.Command {
	return intcli.NewListAdapter(app).Configure()
}
