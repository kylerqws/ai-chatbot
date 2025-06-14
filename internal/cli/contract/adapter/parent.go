package adapter

import "github.com/spf13/cobra"

// ParentAdapter defines the interface for CLI commands with nested subcommands.
type ParentAdapter interface {
	GeneralAdapter

	// Children returns the list of child cobra commands.
	Children() []*cobra.Command

	// ExistChildren returns true if child commands are defined.
	ExistChildren() bool

	// AddChild appends a single child command.
	AddChild(child *cobra.Command)

	// AddChildren appends multiple child commands.
	AddChildren(children ...*cobra.Command)
}
