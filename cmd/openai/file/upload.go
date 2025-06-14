package file

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/openai/file"
)

// UploadCommand creates the OpenAI file upload command.
func UploadCommand(app *intapp.App) *cobra.Command {
	return intcli.NewUploadAdapter(app).Configure()
}
