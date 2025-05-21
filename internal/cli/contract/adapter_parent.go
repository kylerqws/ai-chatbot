package contract

import "github.com/spf13/cobra"

type ParentAdapter interface {
	Adapter

	Children() []*cobra.Command
	ExistChildren() bool

	AddChildren(...*cobra.Command)
	AddChild(*cobra.Command)
}
