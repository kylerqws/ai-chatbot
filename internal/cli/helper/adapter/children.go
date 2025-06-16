package adapter

import "github.com/spf13/cobra"

// ChildrenAdapter provides the implementation for CLI adapter that manage subcommands.
type ChildrenAdapter struct {
	command  *cobra.Command
	children []*cobra.Command
}

// NewChildrenAdapter creates a new instance of ChildrenAdapter.
func NewChildrenAdapter(cmd *cobra.Command) *ChildrenAdapter {
	return &ChildrenAdapter{command: cmd}
}

// Children returns the list of collected child commands.
func (a *ChildrenAdapter) Children() []*cobra.Command {
	return a.children
}

// ExistChildren reports whether any child commands have been added.
func (a *ChildrenAdapter) ExistChildren() bool {
	return len(a.children) > 0
}

// AddChild adds a single child command to the collection.
func (a *ChildrenAdapter) AddChild(child *cobra.Command) {
	if child != nil {
		a.children = append(a.children, child)
	}
}

// AddChildren adds multiple child commands to the collection.
func (a *ChildrenAdapter) AddChildren(children ...*cobra.Command) {
	for i := range children {
		a.AddChild(children[i])
	}
}
