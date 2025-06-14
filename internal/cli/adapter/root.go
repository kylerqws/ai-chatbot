package adapter

import (
	"fmt"
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// RootAdapter implements the root command adapter.
type RootAdapter struct {
	*helper.RootAdapter
}

// NewRootAdapter creates a new root command adapter.
func NewRootAdapter(app *intapp.App) ctr.RootAdapter {
	adp := &RootAdapter{}
	cmd := &cobra.Command{}

	adp.RootAdapter = helper.NewRootAdapter(app, cmd)
	return adp
}

// Configure configures the root CLI command and attaches subcommands.
func (a *RootAdapter) Configure() *cobra.Command {
	app := a.App()

	a.SetUse("chatbot")
	a.SetShort(fmt.Sprintf("CLI for managing %s", app.Name()))
	a.SetVersion(fmt.Sprintf("v%s", app.Version()))

	a.AddChildren(
		action.OpenAICommand(app),
		action.DevCommand(app),
	)

	return a.MainConfigure()
}
