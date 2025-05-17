package adapter

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	"github.com/kylerqws/chatbot/cmd"
	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type RootAdapter struct {
	app      *app.App
	command  *cobra.Command
	children []*cobra.Command
}

func NewRootAdapter(app *app.App) ctr.RootAdapter {
	return &RootAdapter{app: app}
}

func (a *RootAdapter) Configure() *cobra.Command {
	a.command = &cobra.Command{
		Use:     "chatbot",
		Short:   fmt.Sprintf("CLI for managing %s", a.app.Name()),
		Version: fmt.Sprintf("v%s", a.app.Version()),
	}

	a.children = []*cobra.Command{
		cmd.OpenAICommand(a.app),
		cmd.DevCommand(a.app),
	}

	return setup.RootConfigure(a)
}

func (a *RootAdapter) App() *app.App {
	return a.app
}

func (a *RootAdapter) Command() *cobra.Command {
	return a.command
}

func (a *RootAdapter) Children() []*cobra.Command {
	return a.children
}
