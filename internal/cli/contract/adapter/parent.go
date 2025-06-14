package adapter

import "github.com/spf13/cobra"

// ParentAdapter defines the interface for a CLI adapter with nested subcommands.
type ParentAdapter interface {
	GeneralAdapter

	// Children returns all registered child commands.
	Children() []*cobra.Command

	// ExistChildren reports whether any child commands have been added.
	ExistChildren() bool

	// AddChild adds a single child command to the collection.
	AddChild(child *cobra.Command)

	// AddChildren adds multiple child commands to the collection.
	AddChildren(children ...*cobra.Command)
}
