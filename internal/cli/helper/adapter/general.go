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
func (a *GeneralAdapter) App() *app.App {
	return a.app
}

// Command returns the cobra command instance.
func (a *GeneralAdapter) Command() *cobra.Command {
	return a.command
}

// Use returns the usage string.
func (a *GeneralAdapter) Use() string {
	return a.command.Use
}

// SetUse sets the usage string.
func (a *GeneralAdapter) SetUse(value string) {
	a.command.Use = value
}

// Short returns the short description.
func (a *GeneralAdapter) Short() string {
	return a.command.Short
}

// SetShort sets the short description.
func (a *GeneralAdapter) SetShort(value string) {
	a.command.Short = value
}

// Long returns the long description.
func (a *GeneralAdapter) Long() string {
	return a.command.Long
}

// SetLong sets the long description.
func (a *GeneralAdapter) SetLong(value string) {
	a.command.Long = value
}

// Configure applies configuration for the command.
func (a *GeneralAdapter) Configure() *cobra.Command {
	return a.MainConfigure()
}

// MainConfigure applies common configuration for the command.
func (a *GeneralAdapter) MainConfigure() *cobra.Command {
	return setup.GeneralConfigure(a)
}
