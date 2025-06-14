package adapter

import "github.com/spf13/cobra"

// ChildrenAdapter provides base functionality for managing child CLI commands.
type ChildrenAdapter struct {
	command  *cobra.Command
	children []*cobra.Command
}

// NewChildrenAdapter creates a new children command adapter.
func NewChildrenAdapter(cmd *cobra.Command) *ChildrenAdapter {
	return &ChildrenAdapter{command: cmd}
}

// Children returns all registered child commands.
func (h *ChildrenAdapter) Children() []*cobra.Command {
	return h.children
}

// ExistChildren reports whether any child commands have been added.
func (h *ChildrenAdapter) ExistChildren() bool {
	return len(h.children) > 0
}

// AddChild adds a single child command to the collection.
func (h *ChildrenAdapter) AddChild(child *cobra.Command) {
	if child != nil {
		h.children = append(h.children, child)
	}
}

// AddChildren adds multiple child commands to the collection.
func (h *ChildrenAdapter) AddChildren(children ...*cobra.Command) {
	for _, child := range children {
		h.AddChild(child)
	}
}
