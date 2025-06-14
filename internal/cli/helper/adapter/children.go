package adapter

import "github.com/spf13/cobra"

// ChildrenAdapter provides the base implementation for managing CLI subcommands.
type ChildrenAdapter struct {
	command  *cobra.Command
	children []*cobra.Command
}

// NewChildrenAdapter creates a new children command adapter.
func NewChildrenAdapter(cmd *cobra.Command) *ChildrenAdapter {
	return &ChildrenAdapter{command: cmd}
}

// Children returns all registered child commands.
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
	for _, child := range children {
		a.AddChild(child)
	}
}
