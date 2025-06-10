package cmd

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd"
)

// OpenAICommand creates a command for OpenAI-related actions.
func OpenAICommand(app *intapp.App) *cobra.Command {
	return intcli.NewOpenAIAdapter(app).Configure()
}
