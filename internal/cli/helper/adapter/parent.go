package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

// ParentAdapter provides the implementation for CLI adapters with subcommands.
type ParentAdapter struct {
	*GeneralAdapter
	*ChildrenAdapter

	command *cobra.Command
}

// NewParentAdapter creates a new parent command adapter.
func NewParentAdapter(app *app.App, cmd *cobra.Command) *ParentAdapter {
	hlp := &ParentAdapter{command: cmd}

	hlp.GeneralAdapter = NewGeneralAdapter(app, cmd)
	hlp.ChildrenAdapter = NewChildrenAdapter(cmd)

	return hlp
}

// MainConfigure applies common configuration for the command.
func (a *ParentAdapter) MainConfigure() *cobra.Command {
	return setup.ParentConfigure(a)
}
