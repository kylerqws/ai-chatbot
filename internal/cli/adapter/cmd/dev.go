package cmd

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/dev"
	intapp "github.com/kylerqws/chatbot/internal/app"
	hlppar "github.com/kylerqws/chatbot/internal/cli/helper/adapter/parent"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type DevAdapter struct {
	*hlppar.ParentAdapterHelper
}

func NewDevAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &DevAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper =
		hlppar.NewParentAdapterHelper(app, cmd)

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
