package adapter

import "github.com/spf13/cobra"

type ChildrenAdapterHelper struct {
	command  *cobra.Command
	children []*cobra.Command
}

func NewChildrenAdapterHelper(cmd *cobra.Command) *ChildrenAdapterHelper {
	return &ChildrenAdapterHelper{command: cmd}
}

func (h *ChildrenAdapterHelper) Children() []*cobra.Command {
	return h.children
}

func (h *ChildrenAdapterHelper) ExistChildren() bool {
	return len(h.children) > 0
}

func (h *ChildrenAdapterHelper) AddChild(cmd *cobra.Command) {
	h.children = append(h.children, cmd)
}

func (h *ChildrenAdapterHelper) AddChildren(cmds ...*cobra.Command) {
	h.children = append(h.children, cmds...)
}
