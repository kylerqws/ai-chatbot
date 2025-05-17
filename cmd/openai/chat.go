package openai

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	cli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai"
)

func ChatCommand(app *app.App) *cobra.Command {
	return cli.NewJobAdapter(app).Configure()
}
