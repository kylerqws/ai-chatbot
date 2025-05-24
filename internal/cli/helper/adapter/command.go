package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type CommandAdapterHelper struct {
	*GeneralAdapterHelper
	*ErrorAdapterHelper
	*PrintAdapterHelper
	command *cobra.Command
}

func NewCommandAdapterHelper(app *app.App, cmd *cobra.Command) *CommandAdapterHelper {
	hlp := &CommandAdapterHelper{command: cmd}

	hlp.GeneralAdapterHelper = NewGeneralAdapterHelper(app, cmd)
	hlp.ErrorAdapterHelper = NewErrorAdapterHelper(cmd)
	hlp.PrintAdapterHelper = NewPrintAdapterHelper(cmd)

	return hlp
}

func (h *CommandAdapterHelper) FuncArgs() ctr.FuncArgs {
	return ctr.FuncArgs(h.command.Args)
}

func (h *CommandAdapterHelper) SetFuncArgs(fn ctr.FuncArgs) {
	h.command.Args = cobra.PositionalArgs(fn)
}

func (h *CommandAdapterHelper) FuncRunE() ctr.FuncRunE {
	return h.command.RunE
}

func (h *CommandAdapterHelper) SetFuncRunE(fn ctr.FuncRunE) {
	h.command.RunE = fn
}

func (h *CommandAdapterHelper) MainConfigure() *cobra.Command {
	return setup.CommandConfigure(h)
}
