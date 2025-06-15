package adapter

import (
	"github.com/kylerqws/chatbot/internal/app"
	"github.com/spf13/cobra"
)

// GeneralAdapter defines the interface for all CLI adapters.
type GeneralAdapter interface {
	// App returns the associated application instance.
	App() *app.App

	// Command returns the Cobra command instance.
	Command() *cobra.Command

	// Use returns the usage string.
	Use() string

	// SetUse sets the usage string.
	SetUse(value string)

	// Short returns the short description.
	Short() string

	// SetShort sets the short description.
	SetShort(value string)

	// Long returns the long description.
	Long() string

	// SetLong sets the long description.
	SetLong(value string)

	// Configure applies configuration for the command.
	Configure() *cobra.Command

	// MainConfigure applies common configuration for the command.
	MainConfigure() *cobra.Command
}
