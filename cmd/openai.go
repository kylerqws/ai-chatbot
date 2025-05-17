package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	cli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd"
)

func OpenAICommand(app *app.App) *cobra.Command {
	return cli.NewOpenAIAdapter(app).Configure()
}
