package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

// CommandAdapter provides the implementation for CLI adapter with functional command.
type CommandAdapter struct {
	*GeneralAdapter
	*ErrorAdapter
	*PrintAdapter

	command *cobra.Command
}

// NewCommandAdapter creates a new instance of CommandAdapter.
func NewCommandAdapter(app *app.App, cmd *cobra.Command) *CommandAdapter {
	hlp := &CommandAdapter{command: cmd}

	hlp.GeneralAdapter = NewGeneralAdapter(app, cmd)
	hlp.ErrorAdapter = NewErrorAdapter(cmd)
	hlp.PrintAdapter = NewPrintAdapter(cmd)

	return hlp
}

// FuncArgs returns the cobra-compatible argument handler.
func (a *CommandAdapter) FuncArgs() ctr.FuncArgs {
	return ctr.FuncArgs(a.command.Args)
}

// SetFuncArgs sets the cobra-compatible argument handler.
func (a *CommandAdapter) SetFuncArgs(fn ctr.FuncArgs) {
	a.command.Args = cobra.PositionalArgs(fn)
}

// FuncRunE returns the cobra-compatible execution handler.
func (a *CommandAdapter) FuncRunE() ctr.FuncRunE {
	return a.command.RunE
}

// SetFuncRunE sets the cobra-compatible execution handler.
func (a *CommandAdapter) SetFuncRunE(fn ctr.FuncRunE) {
	a.command.RunE = fn
}

// MainConfigure applies common configuration for the command.
func (a *CommandAdapter) MainConfigure() *cobra.Command {
	return setup.CommandConfigure(a)
}
