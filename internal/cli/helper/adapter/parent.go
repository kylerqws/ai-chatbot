package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

type ParentAdapterHelper struct {
	*GeneralAdapterHelper
	*ChildrenAdapterHelper
	command *cobra.Command
}

func NewParentAdapterHelper(app *app.App, cmd *cobra.Command) *ParentAdapterHelper {
	hlp := &ParentAdapterHelper{command: cmd}

	hlp.GeneralAdapterHelper = NewGeneralAdapterHelper(app, cmd)
	hlp.ChildrenAdapterHelper = NewChildrenAdapterHelper(cmd)

	return hlp
}

func (h *ParentAdapterHelper) MainConfigure() *cobra.Command {
	return setup.ParentConfigure(h)
}
