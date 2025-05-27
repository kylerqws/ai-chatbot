package adapter

import "github.com/spf13/cobra"

type ChildrenAdapter struct {
	command  *cobra.Command
	children []*cobra.Command
}

func NewChildrenAdapter(cmd *cobra.Command) *ChildrenAdapter {
	return &ChildrenAdapter{command: cmd}
}

func (h *ChildrenAdapter) Children() []*cobra.Command {
	return h.children
}

func (h *ChildrenAdapter) ExistChildren() bool {
	return len(h.children) > 0
}

func (h *ChildrenAdapter) AddChild(cmd *cobra.Command) {
	if cmd != nil {
		h.children = append(h.children, cmd)
	}
}

func (h *ChildrenAdapter) AddChildren(cmds ...*cobra.Command) {
	for i := range cmds {
		h.AddChild(cmds[i])
	}
}
