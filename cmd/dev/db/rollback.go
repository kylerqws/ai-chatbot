package db

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	cli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/dev/db"
)

func RollbackCommand(app *app.App) *cobra.Command {
	return cli.NewRollbackAdapter(app).Configure()
}
