package openai

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai"
)

// ChatCommand creates the OpenAI chat command.
func ChatCommand(app *intapp.App) *cobra.Command {
	return intcli.NewChatAdapter(app).Configure()
}
