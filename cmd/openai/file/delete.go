package file

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai/file"
)

// DeleteCommand creates a command for deleting an OpenAI file.
func DeleteCommand(app *intapp.App) *cobra.Command {
	return intcli.NewDeleteAdapter(app).Configure()
}
