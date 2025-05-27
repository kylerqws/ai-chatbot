package adapter

import (
	"github.com/kylerqws/chatbot/internal/cli/setup"
	"github.com/spf13/cobra"
)

type FlagAdapter struct {
	command *cobra.Command
}

func NewFlagAdapter(cmd *cobra.Command) *FlagAdapter {
	return &FlagAdapter{command: cmd}
}

func (h *FlagAdapter) AddStringFlag(name, shorthand, value, desc string) {
	setup.AddStringFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapter) AddStringSliceFlag(name, shorthand string, value []string, desc string) {
	setup.AddStringSliceFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapter) AddBoolFlag(name, shorthand string, value bool, desc string) {
	setup.AddBoolFlag(h.command, name, shorthand, value, desc, false)
}
