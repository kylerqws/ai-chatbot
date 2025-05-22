package cmd

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/dev"
	intapp "github.com/kylerqws/chatbot/internal/app"
	inthlp "github.com/kylerqws/chatbot/internal/cli/helper"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type DevAdapter struct {
	*inthlp.ParentAdapterHelper
}

func NewDevAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &DevAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper =
		inthlp.NewParentAdapterHelper(app, cmd)

	return adp
}

func (a *DevAdapter) Configure() *cobra.Command {
	app := a.App()

	a.SetUse("dev")
	a.SetShort("Tools for application development")

	a.AddChildren(
		action.DBCommand(app),
	)

	return a.MainConfigure()
}
