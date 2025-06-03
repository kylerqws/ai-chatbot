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

func (h *FlagAdapter) AddUint8Flag(name, shorthand string, value uint8, desc string) {
	setup.AddUint8Flag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapter) AddUint8SliceFlag(name, shorthand string, value []uint8, desc string) {
	setup.AddUint8SliceFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapter) AddUintFlag(name, shorthand string, value uint, desc string) {
	setup.AddUintFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapter) AddUintSliceFlag(name, shorthand string, value []uint, desc string) {
	setup.AddUintSliceFlag(h.command, name, shorthand, value, desc, false)
}

func (h *FlagAdapter) AddBoolFlag(name, shorthand string, value bool, desc string) {
	setup.AddBoolFlag(h.command, name, shorthand, value, desc, false)
}
