package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

type GeneralAdapterHelper struct {
	app     *app.App
	command *cobra.Command
}

func NewGeneralAdapterHelper(app *app.App, cmd *cobra.Command) *GeneralAdapterHelper {
	return &GeneralAdapterHelper{app: app, command: cmd}
}

func (h *GeneralAdapterHelper) App() *app.App {
	return h.app
}

func (h *GeneralAdapterHelper) Command() *cobra.Command {
	return h.command
}

func (h *GeneralAdapterHelper) Use() string {
	return h.command.Use
}

func (h *GeneralAdapterHelper) SetUse(use string) {
	h.command.Use = use
}

func (h *GeneralAdapterHelper) Short() string {
	return h.command.Short
}

func (h *GeneralAdapterHelper) SetShort(short string) {
	h.command.Short = short
}

func (h *GeneralAdapterHelper) Configure() *cobra.Command {
	return h.MainConfigure()
}

func (h *GeneralAdapterHelper) MainConfigure() *cobra.Command {
	return setup.GeneralConfigure(h)
}
