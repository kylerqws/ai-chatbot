package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

// RootAdapter provides the base implementation for the root CLI adapter.
type RootAdapter struct {
	*ParentAdapter
	command *cobra.Command
}

// NewRootAdapter creates a new root adapter.
func NewRootAdapter(app *app.App, cmd *cobra.Command) *RootAdapter {
	hlp := &RootAdapter{command: cmd}
	hlp.ParentAdapter = NewParentAdapter(app, cmd)

	return hlp
}

// Version returns the CLI application version.
func (h *RootAdapter) Version() string {
	return h.command.Version
}

// SetVersion sets the CLI application version.
func (h *RootAdapter) SetVersion(version string) {
	h.command.Version = version
}

// MainConfigure applies common configuration for the command.
func (h *RootAdapter) MainConfigure() *cobra.Command {
	return setup.RootConfigure(h)
}
