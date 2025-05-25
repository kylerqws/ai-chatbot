package adapter

import (
	"github.com/kylerqws/chatbot/internal/cli/setup"
	"github.com/spf13/cobra"
)

type FlagAdapterHelper struct {
	command *cobra.Command
	errors  []error
}

func NewFlagAdapterHelper(cmd *cobra.Command) *FlagAdapterHelper {
	return &FlagAdapterHelper{command: cmd}
}

func (h *FlagAdapterHelper) AddStringFlag(name, shorthand, value, desc string) {
	setup.AddStringFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapterHelper) AddStringSliceFlag(name, shorthand string, value []string, desc string) {
	setup.AddStringSliceFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapterHelper) AddBoolFlag(name, shorthand string, value bool, desc string) {
	setup.AddBoolFlag(h.command, name, shorthand, value, desc, false)
}
