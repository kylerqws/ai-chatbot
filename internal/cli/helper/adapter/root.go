package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

type RootAdapter struct {
	*ParentAdapter
	command *cobra.Command
}

func NewRootAdapter(app *app.App, cmd *cobra.Command) *RootAdapter {
	hlp := &RootAdapter{command: cmd}
	hlp.ParentAdapter = NewParentAdapter(app, cmd)

	return hlp
}

func (h *RootAdapter) Version() string {
	return h.command.Version
}

func (h *RootAdapter) SetVersion(version string) {
	h.command.Version = version
}

func (h *RootAdapter) MainConfigure() *cobra.Command {
	return setup.RootConfigure(h)
}
