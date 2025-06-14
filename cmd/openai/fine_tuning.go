package openai

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai"
)

// FineTuningCommand creates the OpenAI fine-tuning command.
func FineTuningCommand(app *intapp.App) *cobra.Command {
	return intcli.NewFineTuningAdapter(app).Configure()
}
