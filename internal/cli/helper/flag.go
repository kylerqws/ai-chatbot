package helper

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/cli/setup/flag"
)

type FlagAdapterHelper struct {
	command *cobra.Command
	errors  []error
}

func NewFlagAdapterHelper(cmd *cobra.Command) *FlagAdapterHelper {
	return &FlagAdapterHelper{command: cmd}
}

func (h *FlagAdapterHelper) AddStringFlag(name, shorthand, value, desc string) {
	flag.AddStringFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapterHelper) AddStringSliceFlag(name, shorthand string, value []string, desc string) {
	flag.AddStringSliceFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapterHelper) AddInt64Flag(name, shorthand string, value int64, desc string) {
	flag.AddInt64Flag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapterHelper) AddInt64SliceFlag(name, shorthand string, value []int64, desc string) {
	flag.AddInt64SliceFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapterHelper) AddBoolFlag(name, shorthand string, value bool, desc string) {
	flag.AddBoolFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapterHelper) HasAnyFlag(names ...string) bool {
	fgs := h.command.Flags()
	for i := range names {
		if fgs.Changed(names[i]) {
			return true
		}
	}
	return false
}
