package openai

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai"
)

// FileCommand creates the OpenAI file command.
func FileCommand(app *intapp.App) *cobra.Command {
	return intcli.NewFileAdapter(app).Configure()
}
