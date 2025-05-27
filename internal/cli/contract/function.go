package contract

import "github.com/spf13/cobra"

type (
	FuncArgs func(*cobra.Command, []string) error
	FuncRunE func(*cobra.Command, []string) error
	FuncHelp func(*cobra.Command, []string)

	FuncValidateString func(string) error
)
