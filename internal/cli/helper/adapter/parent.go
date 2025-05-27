package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

type ParentAdapter struct {
	*GeneralAdapter
	*ChildrenAdapter
	command *cobra.Command
}

func NewParentAdapter(app *app.App, cmd *cobra.Command) *ParentAdapter {
	hlp := &ParentAdapter{command: cmd}

	hlp.GeneralAdapter = NewGeneralAdapter(app, cmd)
	hlp.ChildrenAdapter = NewChildrenAdapter(cmd)

	return hlp
}

func (h *ParentAdapter) MainConfigure() *cobra.Command {
	return setup.ParentConfigure(h)
}
