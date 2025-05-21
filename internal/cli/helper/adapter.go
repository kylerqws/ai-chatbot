package helper

import (
	intapp "github.com/kylerqws/chatbot/internal/app"
	"github.com/spf13/cobra"
)

type AdapterHelper struct {
	app     *intapp.App
	command *cobra.Command
}

func NewAdapterHelper(app *intapp.App, cmd *cobra.Command) *AdapterHelper {
	return &AdapterHelper{app: app, command: cmd}
}

func (h *AdapterHelper) App() *intapp.App {
	return h.app
}

func (h *AdapterHelper) Command() *cobra.Command {
	return h.command
}

func (h *AdapterHelper) Use() string {
	return h.command.Use
}

func (h *AdapterHelper) SetUse(use string) {
	h.command.Use = use
}

func (h *AdapterHelper) Short() string {
	return h.command.Short
}

func (h *AdapterHelper) SetShort(short string) {
	h.command.Short = short
}

func (h *AdapterHelper) Configure() *cobra.Command {
	return h.MainConfigure()
}

func (h *AdapterHelper) MainConfigure() *cobra.Command {
	return h.command
}
