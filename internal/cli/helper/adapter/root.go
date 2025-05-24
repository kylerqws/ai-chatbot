package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

type RootAdapterHelper struct {
	*ParentAdapterHelper
	command *cobra.Command
}

func NewRootAdapterHelper(app *app.App, cmd *cobra.Command) *RootAdapterHelper {
	hlp := &RootAdapterHelper{command: cmd}
	hlp.ParentAdapterHelper = NewParentAdapterHelper(app, cmd)

	return hlp
}

func (h *RootAdapterHelper) Version() string {
	return h.command.Version
}

func (h *RootAdapterHelper) SetVersion(version string) {
	h.command.Version = version
}

func (h *RootAdapterHelper) MainConfigure() *cobra.Command {
	return setup.RootConfigure(h)
}
