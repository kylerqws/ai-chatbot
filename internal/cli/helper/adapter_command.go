package helper

import (
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type CommandAdapterHelper struct {
	*AdapterHelper
	*ErrorAdapterHelper
	*PrintAdapterHelper
}

func NewCommandAdapterHelper(app *intapp.App, cmd *cobra.Command) *CommandAdapterHelper {
	hlp := &CommandAdapterHelper{}

	hlp.AdapterHelper = NewAdapterHelper(app, cmd)
	hlp.ErrorAdapterHelper = NewErrorAdapterHelper(cmd)
	hlp.PrintAdapterHelper = NewPrintAdapterHelper(cmd)

	return hlp
}

func (h *CommandAdapterHelper) FuncArgs() ctr.FuncArgs {
	return ctr.FuncArgs(h.AdapterHelper.command.Args)
}

func (h *CommandAdapterHelper) SetFuncArgs(fn ctr.FuncArgs) {
	h.AdapterHelper.command.Args = cobra.PositionalArgs(fn)
}

func (h *CommandAdapterHelper) FuncRunE() ctr.FuncRunE {
	return h.AdapterHelper.command.RunE
}

func (h *CommandAdapterHelper) SetFuncRunE(fn ctr.FuncRunE) {
	h.AdapterHelper.command.RunE = fn
}

func (h *CommandAdapterHelper) MainConfigure() *cobra.Command {
	return setup.CommandConfigure(h)
}
