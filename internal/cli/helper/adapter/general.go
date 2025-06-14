package adapter

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"
)

// GeneralAdapter provides the base implementation for CLI adapters.
type GeneralAdapter struct {
	app     *app.App
	command *cobra.Command
}

// NewGeneralAdapter creates a new general command adapter.
func NewGeneralAdapter(app *app.App, cmd *cobra.Command) *GeneralAdapter {
	return &GeneralAdapter{app: app, command: cmd}
}

// App returns the associated application instance.
func (h *GeneralAdapter) App() *app.App {
	return h.app
}

// Command returns the cobra command instance.
func (h *GeneralAdapter) Command() *cobra.Command {
	return h.command
}

// Use returns the usage string.
func (h *GeneralAdapter) Use() string {
	return h.command.Use
}

// SetUse sets the usage string.
func (h *GeneralAdapter) SetUse(value string) {
	h.command.Use = value
}

// Short returns the short description.
func (h *GeneralAdapter) Short() string {
	return h.command.Short
}

// SetShort sets the short description.
func (h *GeneralAdapter) SetShort(value string) {
	h.command.Short = value
}

// Long returns the long description.
func (h *GeneralAdapter) Long() string {
	return h.command.Long
}

// SetLong sets the long description.
func (h *GeneralAdapter) SetLong(value string) {
	h.command.Long = value
}

// Configure applies full configuration for the command.
func (h *GeneralAdapter) Configure() *cobra.Command {
	return h.MainConfigure()
}

// MainConfigure applies common configuration for the command.
func (h *GeneralAdapter) MainConfigure() *cobra.Command {
	return setup.GeneralConfigure(h)
}
