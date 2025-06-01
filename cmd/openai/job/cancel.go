package job

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai/job"
)

func CancelCommand(app *intapp.App) *cobra.Command {
	return intcli.NewCancelAdapter(app).Configure()
}
