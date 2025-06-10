package openai

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai"
)

// ChatCommand creates a command for interacting with the OpenAI chat API.
func ChatCommand(app *intapp.App) *cobra.Command {
	return intcli.NewChatAdapter(app).Configure()
}
