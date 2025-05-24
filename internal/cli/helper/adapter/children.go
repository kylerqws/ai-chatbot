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
	if cmd != nil {
		h.children = append(h.children, cmd)
	}
}

func (h *ChildrenAdapterHelper) AddChildren(cmds ...*cobra.Command) {
	for i := range cmds {
		if cmds[i] != nil {
			h.children = append(h.children, cmds[i])
		}
	}
}
