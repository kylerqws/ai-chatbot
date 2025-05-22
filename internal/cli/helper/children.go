package helper

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

func (h *ChildrenAdapterHelper) AddChildren(cmd ...*cobra.Command) {
	h.children = append(h.children, cmd...)
}

func (h *ChildrenAdapterHelper) AddChild(cmd *cobra.Command) {
	h.AddChildren(cmd)
}

func (h *ChildrenAdapterHelper) ExistChildren() bool {
	return len(h.children) > 0
}
