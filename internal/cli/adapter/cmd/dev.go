package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	"github.com/kylerqws/chatbot/cmd/dev"
	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type DevAdapter struct {
	app      *app.App
	command  *cobra.Command
	children []*cobra.Command
}

func NewDevAdapter(app *app.App) ctr.ParentAdapter {
	return &DevAdapter{app: app}
}

func (a *DevAdapter) Configure() *cobra.Command {
	a.command = &cobra.Command{
		Use:   "dev",
		Short: "Tools for application development",
	}

	a.children = []*cobra.Command{
		dev.DBCommand(a.app),
	}

	return setup.ParentConfigure(a)
}

func (a *DevAdapter) App() *app.App {
	return a.app
}

func (a *DevAdapter) Command() *cobra.Command {
	return a.command
}

func (a *DevAdapter) Children() []*cobra.Command {
	return a.children
}
