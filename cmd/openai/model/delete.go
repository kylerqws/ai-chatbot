package file

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai/model"
)

// DeleteCommand creates a command for deleting an OpenAI model.
func DeleteCommand(app *intapp.App) *cobra.Command {
	return intcli.NewDeleteAdapter(app).Configure()
}
