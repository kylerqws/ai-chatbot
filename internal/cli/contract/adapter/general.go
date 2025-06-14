package adapter

import (
	"github.com/kylerqws/chatbot/internal/app"
	"github.com/spf13/cobra"
)

// GeneralAdapter defines the interface for all CLI command adapters.
type GeneralAdapter interface {
	// App returns the associated application instance.
	App() *app.App

	// Command returns the configured cobra command.
	Command() *cobra.Command

	// Use returns the usage string.
	Use() string

	// SetUse sets the usage string.
	SetUse(value string)

	// Short returns the short description for the command.
	Short() string

	// SetShort sets the short description.
	SetShort(value string)

	// Long returns the long description for the command.
	Long() string

	// SetLong sets the long description.
	SetLong(value string)

	// Configure performs command-specific configuration logic.
	Configure() *cobra.Command

	// MainConfigure performs the common configuration for all commands.
	MainConfigure() *cobra.Command
}
