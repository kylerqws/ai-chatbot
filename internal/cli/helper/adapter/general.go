package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

type GeneralAdapter struct {
	app     *app.App
	command *cobra.Command
}

func NewGeneralAdapter(app *app.App, cmd *cobra.Command) *GeneralAdapter {
	return &GeneralAdapter{app: app, command: cmd}
}

func (h *GeneralAdapter) App() *app.App {
	return h.app
}

func (h *GeneralAdapter) Command() *cobra.Command {
	return h.command
}

func (h *GeneralAdapter) Use() string {
	return h.command.Use
}

func (h *GeneralAdapter) SetUse(use string) {
	h.command.Use = use
}

func (h *GeneralAdapter) Short() string {
	return h.command.Short
}

func (h *GeneralAdapter) SetShort(short string) {
	h.command.Short = short
}

func (h *GeneralAdapter) Configure() *cobra.Command {
	return h.MainConfigure()
}

func (h *GeneralAdapter) MainConfigure() *cobra.Command {
	return setup.GeneralConfigure(h)
}
