package dev

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intcli "github.com/kylerqws/chatbot/internal/cli/adapter/cmd/dev"
)

func DBCommand(app *intapp.App) *cobra.Command {
	return intcli.NewDBAdapter(app).Configure()
}
