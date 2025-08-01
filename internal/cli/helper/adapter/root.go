package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

// RootAdapter provides the implementation for the root CLI adapter.
type RootAdapter struct {
	*ParentAdapter
	command *cobra.Command
}

// NewRootAdapter creates a new instance of RootAdapter.
func NewRootAdapter(app *app.App, cmd *cobra.Command) *RootAdapter {
	hlp := &RootAdapter{command: cmd}
	hlp.ParentAdapter = NewParentAdapter(app, cmd)

	return hlp
}

// Version returns the CLI application version.
func (a *RootAdapter) Version() string {
	return a.command.Version
}

// SetVersion sets the CLI application version.
func (a *RootAdapter) SetVersion(version string) {
	a.command.Version = version
}

// MainConfigure applies common configuration for the command.
func (a *RootAdapter) MainConfigure() *cobra.Command {
	return setup.RootConfigure(a)
}
