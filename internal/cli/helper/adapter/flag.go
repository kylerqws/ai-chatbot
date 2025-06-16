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

// HasChangedFlagInCache checks if the flag was previously marked as changed.
func (a *FlagAdapter) HasChangedFlagInCache(key string) bool {
	if a.changedFlags == nil {
		return false
	}
	return a.changedFlags[key]
}

// CacheChangedFlag checks and caches whether the specified flag was changed.
func (a *FlagAdapter) CacheChangedFlag(key string) {
	if _, ok := a.changedFlags[key]; !ok {
		a.changedFlags[key] = a.command.Flags().Changed(key)
	}
}

// HasChangedFlagsInCache checks if any flag was marked as changed.
func (a *FlagAdapter) HasChangedFlagsInCache() bool {
	for key := range a.changedFlags {
		if a.HasChangedFlagInCache(key) {
			return true
		}
	}
	return false
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

// GetStringFlag retrieves the value of a string flag.
func (a *FlagAdapter) GetStringFlag(key string) (string, error) {
	return a.command.Flags().GetString(key)
}

// GetPointerStringFlag retrieves the value of a string flag as pointer.
func (a *FlagAdapter) GetPointerStringFlag(key string) (*string, error) {
	val, err := a.GetStringFlag(key)
	return &val, err
}

// AddStringSliceFlag adds a string slice flag to the command.
func (a *FlagAdapter) AddStringSliceFlag(name, shorthand string, value []string, desc string) {
	setup.AddStringSliceFlag(a.command, name, shorthand, value, desc, false)
}

// GetStringSliceFlag retrieves the value of a string slice flag.
func (a *FlagAdapter) GetStringSliceFlag(key string) ([]string, error) {
	return a.command.Flags().GetStringSlice(key)
}

// GetPointerStringSliceFlag retrieves the value of a string slice flag as pointer.
func (a *FlagAdapter) GetPointerStringSliceFlag(key string) (*[]string, error) {
	val, err := a.GetStringSliceFlag(key)
	return &val, err
}

// AddUint8Flag adds an uint8 flag to the command.
func (a *FlagAdapter) AddUint8Flag(name, shorthand string, value uint8, desc string) {
	setup.AddUint8Flag(a.command, name, shorthand, value, desc, false)
}

// GetUint8Flag retrieves the value of an uint8 flag.
func (a *FlagAdapter) GetUint8Flag(key string) (uint8, error) {
	return a.command.Flags().GetUint8(key)
}

// GetPointerUint8Flag retrieves the value of an uint8 flag as pointer.
func (a *FlagAdapter) GetPointerUint8Flag(key string) (*uint8, error) {
	val, err := a.GetUint8Flag(key)
	return &val, err
}

// AddUintFlag adds an uint flag to the command.
func (a *FlagAdapter) AddUintFlag(name, shorthand string, value uint, desc string) {
	setup.AddUintFlag(a.command, name, shorthand, value, desc, false)
}

// GetUintFlag retrieves the value of an uint flag.
func (a *FlagAdapter) GetUintFlag(key string) (uint, error) {
	return a.command.Flags().GetUint(key)
}

// GetPointerUintFlag retrieves the value of an uint flag as pointer.
func (a *FlagAdapter) GetPointerUintFlag(key string) (*uint, error) {
	val, err := a.GetUintFlag(key)
	return &val, err
}

// AddUintSliceFlag adds an uint slice flag to the command.
func (a *FlagAdapter) AddUintSliceFlag(name, shorthand string, value []uint, desc string) {
	setup.AddUintSliceFlag(a.command, name, shorthand, value, desc, false)
}

// GetUintSliceFlag retrieves the value of an uint slice flag.
func (a *FlagAdapter) GetUintSliceFlag(key string) ([]uint, error) {
	return a.command.Flags().GetUintSlice(key)
}

// GetPointerUintSliceFlag retrieves the value of a uint slice flag as pointer.
func (a *FlagAdapter) GetPointerUintSliceFlag(key string) (*[]uint, error) {
	val, err := a.GetUintSliceFlag(key)
	return &val, err
}

// AddBoolFlag adds a bool flag to the command.
func (a *FlagAdapter) AddBoolFlag(name, shorthand string, value bool, desc string) {
	setup.AddBoolFlag(a.command, name, shorthand, value, desc, false)
}

// GetBoolFlag retrieves the value of a bool flag.
func (a *FlagAdapter) GetBoolFlag(key string) (bool, error) {
	return a.command.Flags().GetBool(key)
}

// GetPointerBoolFlag retrieves the value of a bool flag as pointer.
func (a *FlagAdapter) GetPointerBoolFlag(key string) (*bool, error) {
	val, err := a.GetBoolFlag(key)
	return &val, err
}
