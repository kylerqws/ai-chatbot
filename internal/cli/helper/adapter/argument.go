package adapter

import (
	"github.com/spf13/cobra"
	"strings"
)

const SeparatorArgParts = ":"

type ArgumentAdapter struct {
	command *cobra.Command
}

func NewArgumentAdapter(cmd *cobra.Command) *ArgumentAdapter {
	return &ArgumentAdapter{command: cmd}
}

func (h *ArgumentAdapter) SplitDoubleArg(arg string) (string, string) {
	parts := strings.SplitN(arg, SeparatorArgParts, 2)
	parts = h.trimSpace(parts)

	switch len(parts) {
	case 2:
		return parts[0], parts[1]
	case 1:
		return parts[0], ""
	default:
		return "", ""
	}
}

func (h *ArgumentAdapter) SplitTripleArg(arg string) (string, string, string) {
	parts := strings.SplitN(arg, SeparatorArgParts, 3)
	parts = h.trimSpace(parts)

	switch len(parts) {
	case 3:
		return parts[0], parts[1], parts[2]
	case 2:
		return parts[0], parts[1], ""
	case 1:
		return parts[0], "", ""
	default:
		return "", "", ""
	}
}

func (*ArgumentAdapter) trimSpace(vals []string) []string {
	for i := range vals {
		vals[i] = strings.TrimSpace(vals[i])
	}
	return vals
}
