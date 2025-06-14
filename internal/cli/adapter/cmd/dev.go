package cmd

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/dev"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// DevAdapter provides the implementation for the dev CLI adapter.
type DevAdapter struct {
	*helper.ParentAdapter
}

// NewDevAdapter creates a new dev command adapter.
func NewDevAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &DevAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapter = helper.NewParentAdapter(app, cmd)
	return adp
}

// Configure applies configuration for the command.
func (a *DevAdapter) Configure() *cobra.Command {
	app := a.App()

	a.SetUse("dev")
	a.SetShort("Tools for application development")

	a.AddChildren(
		action.DBCommand(app),
	)

	return a.MainConfigure()
}
