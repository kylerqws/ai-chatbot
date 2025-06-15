package contract

import "github.com/spf13/cobra"

type (
	// FuncArgs defines the handler for processing arguments passed to a command.
	FuncArgs func(cmd *cobra.Command, args []string) error

	// FuncRunE defines the main logic handler for executing a command.
	FuncRunE func(cmd *cobra.Command, args []string) error

	// FuncHelp defines a handler for rendering custom help output.
	FuncHelp func(cmd *cobra.Command, args []string)

	// FuncValidateString defines a validator for a single string value.
	FuncValidateString func(value string) error

	// FuncValidateUint8 defines a validator for a single uint8 value.
	FuncValidateUint8 func(value uint8) error

	// FuncValidateUint defines a validator for a single uint value.
	FuncValidateUint func(value uint) error
)
