package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type CommandAdapter struct {
	*GeneralAdapter
	*ErrorAdapter
	*PrintAdapter
	command *cobra.Command
}

func NewCommandAdapter(app *app.App, cmd *cobra.Command) *CommandAdapter {
	hlp := &CommandAdapter{command: cmd}

	hlp.GeneralAdapter = NewGeneralAdapter(app, cmd)
	hlp.ErrorAdapter = NewErrorAdapter(cmd)
	hlp.PrintAdapter = NewPrintAdapter(cmd)

	return hlp
}

func (h *CommandAdapter) FuncArgs() ctr.FuncArgs {
	return ctr.FuncArgs(h.command.Args)
}

func (h *CommandAdapter) SetFuncArgs(fn ctr.FuncArgs) {
	h.command.Args = cobra.PositionalArgs(fn)
}

func (h *CommandAdapter) FuncRunE() ctr.FuncRunE {
	return h.command.RunE
}

func (h *CommandAdapter) SetFuncRunE(fn ctr.FuncRunE) {
	h.command.RunE = fn
}

func (h *CommandAdapter) MainConfigure() *cobra.Command {
	return setup.CommandConfigure(h)
}
