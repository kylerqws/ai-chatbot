package helper

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

type RootAdapterHelper struct {
	*ParentAdapterHelper
}

func NewRootAdapterHelper(app *intapp.App, cmd *cobra.Command) *RootAdapterHelper {
	hlp := &RootAdapterHelper{}
	hlp.ParentAdapterHelper = NewParentAdapterHelper(app, cmd)

	return hlp
}

func (h *RootAdapterHelper) Version() string {
	return h.ParentAdapterHelper.AdapterHelper.command.Version
}

func (h *RootAdapterHelper) SetVersion(version string) {
	h.ParentAdapterHelper.AdapterHelper.command.Version = version
}

func (h *RootAdapterHelper) MainConfigure() *cobra.Command {
	return setup.RootConfigure(h)
}
