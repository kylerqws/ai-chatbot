package openai

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai"
)

// ModelCommand creates the OpenAI model command.
func ModelCommand(app *intapp.App) *cobra.Command {
	return intcli.NewModelAdapter(app).Configure()
}
