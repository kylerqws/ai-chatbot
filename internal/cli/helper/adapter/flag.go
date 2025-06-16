package adapter

import (
	"github.com/kylerqws/chatbot/internal/cli/setup"
	"github.com/spf13/cobra"
)

// FlagAdapter provides the implementation for CLI adapter with flag handling.
type FlagAdapter struct {
	command      *cobra.Command
	changedFlags map[string]bool
}

// NewFlagAdapter creates a new instance of FlagAdapter.
func NewFlagAdapter(cmd *cobra.Command) *FlagAdapter {
	return &FlagAdapter{command: cmd, changedFlags: map[string]bool{}}
}

// HasChangedFlag checks if the flag was previously marked as changed.
func (a *FlagAdapter) HasChangedFlag(key string) bool {
	if a.changedFlags == nil {
		return false
	}
	return a.changedFlags[key]
}

// HasChangedFlags checks if any flags was previously marked as changed.
func (a *FlagAdapter) HasChangedFlags() bool {
	for key := range a.changedFlags {
		if a.HasChangedFlag(key) {
			return true
		}
	}
	return false
}

// CacheChangedFlag checks and caches whether the specified flag was changed.
func (a *FlagAdapter) CacheChangedFlag(key string) {
	if _, ok := a.changedFlags[key]; !ok {
		a.changedFlags[key] = a.command.Flags().Changed(key)
	}
}

// CacheChangedFlags checks and caches whether the specified flags were changed.
func (a *FlagAdapter) CacheChangedFlags(keys ...string) {
	for i := range keys {
		a.CacheChangedFlag(keys[i])
	}
}

// AddStringFlag adds a string flag to the command.
func (a *FlagAdapter) AddStringFlag(name, shorthand, value, desc string) {
	setup.AddStringFlag(a.command, name, shorthand, value, desc, false)
}

// StringFlag retrieves the value of a string flag.
func (a *FlagAdapter) StringFlag(key string) (string, error) {
	return a.command.Flags().GetString(key)
}

// PointerStringFlag retrieves the value of a string flag as pointer.
func (a *FlagAdapter) PointerStringFlag(key string) (*string, error) {
	val, err := a.StringFlag(key)
	return &val, err
}

// AddStringSliceFlag adds a string slice flag to the command.
func (a *FlagAdapter) AddStringSliceFlag(name, shorthand string, value []string, desc string) {
	setup.AddStringSliceFlag(a.command, name, shorthand, value, desc, false)
}

// StringSliceFlag retrieves the value of a string slice flag.
func (a *FlagAdapter) StringSliceFlag(key string) ([]string, error) {
	return a.command.Flags().GetStringSlice(key)
}

// PointerStringSliceFlag retrieves the value of a string slice flag as pointer.
func (a *FlagAdapter) PointerStringSliceFlag(key string) (*[]string, error) {
	val, err := a.StringSliceFlag(key)
	return &val, err
}

// AddUint8Flag adds an uint8 flag to the command.
func (a *FlagAdapter) AddUint8Flag(name, shorthand string, value uint8, desc string) {
	setup.AddUint8Flag(a.command, name, shorthand, value, desc, false)
}

// Uint8Flag retrieves the value of an uint8 flag.
func (a *FlagAdapter) Uint8Flag(key string) (uint8, error) {
	return a.command.Flags().GetUint8(key)
}

// PointerUint8Flag retrieves the value of an uint8 flag as pointer.
func (a *FlagAdapter) PointerUint8Flag(key string) (*uint8, error) {
	val, err := a.Uint8Flag(key)
	return &val, err
}

// AddUintFlag adds an uint flag to the command.
func (a *FlagAdapter) AddUintFlag(name, shorthand string, value uint, desc string) {
	setup.AddUintFlag(a.command, name, shorthand, value, desc, false)
}

// UintFlag retrieves the value of an uint flag.
func (a *FlagAdapter) UintFlag(key string) (uint, error) {
	return a.command.Flags().GetUint(key)
}

// PointerUintFlag retrieves the value of an uint flag as pointer.
func (a *FlagAdapter) PointerUintFlag(key string) (*uint, error) {
	val, err := a.UintFlag(key)
	return &val, err
}

// AddUintSliceFlag adds an uint slice flag to the command.
func (a *FlagAdapter) AddUintSliceFlag(name, shorthand string, value []uint, desc string) {
	setup.AddUintSliceFlag(a.command, name, shorthand, value, desc, false)
}

// UintSliceFlag retrieves the value of an uint slice flag.
func (a *FlagAdapter) UintSliceFlag(key string) ([]uint, error) {
	return a.command.Flags().GetUintSlice(key)
}

// PointerUintSliceFlag retrieves the value of an uint slice flag as pointer.
func (a *FlagAdapter) PointerUintSliceFlag(key string) (*[]uint, error) {
	val, err := a.UintSliceFlag(key)
	return &val, err
}

// AddBoolFlag adds a bool flag to the command.
func (a *FlagAdapter) AddBoolFlag(name, shorthand string, value bool, desc string) {
	setup.AddBoolFlag(a.command, name, shorthand, value, desc, false)
}

// BoolFlag retrieves the value of a bool flag.
func (a *FlagAdapter) BoolFlag(key string) (bool, error) {
	return a.command.Flags().GetBool(key)
}

// PointerBoolFlag retrieves the value of a bool flag as pointer.
func (a *FlagAdapter) PointerBoolFlag(key string) (*bool, error) {
	val, err := a.BoolFlag(key)
	return &val, err
}
