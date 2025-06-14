package contract

import "github.com/spf13/cobra"

type (
	// FuncArgs is a command handler that accepts arguments.
	FuncArgs func(cmd *cobra.Command, args []string) error

	// FuncRunE is a cobra-compatible command executor.
	FuncRunE func(cmd *cobra.Command, args []string) error

	// FuncHelp renders help content for the command.
	FuncHelp func(cmd *cobra.Command, args []string)

	// FuncValidateString checks the validity of a string value.
	FuncValidateString func(value string) error

	// FuncValidateUint8 checks the validity of an uint8 value.
	FuncValidateUint8 func(value uint8) error

	// FuncValidateUint checks the validity of an uint value.
	FuncValidateUint func(value uint) error
)
