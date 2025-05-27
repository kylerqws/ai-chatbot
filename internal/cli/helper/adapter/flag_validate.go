package adapter

import (
	"fmt"
	"github.com/spf13/cobra"
)

type ValidateFlagAdapter struct {
	*FlagAdapter
	command *cobra.Command
}

func NewValidateFlagAdapter(cmd *cobra.Command) *ValidateFlagAdapter {
	hlp := &ValidateFlagAdapter{command: cmd}
	hlp.FlagAdapter = NewFlagAdapter(cmd)

	return hlp
}

func (h *FlagAdapter) HasAnyFlag(names ...string) bool {
	fgs := h.command.Flags()
	for i := range names {
		if fgs.Changed(names[i]) {
			return true
		}
	}
	return false
}

func (h *ValidateFlagAdapter) HasMoreArgsThan(count uint8) bool {
	return len(h.command.Flags().Args()) > int(count)
}

func (h *ValidateFlagAdapter) IsFlagChanged(flagKey string) bool {
	return h.command.Flags().Changed(flagKey)
}

func (h *ValidateFlagAdapter) IsFlagRequired(flagKey string) bool {
	f := h.command.Flags().Lookup(flagKey)
	return f != nil && f.Annotations != nil && f.Annotations[cobra.BashCompOneRequiredFlag] != nil
}

func (h *ValidateFlagAdapter) ValidateHasAnyFlags(flagKeys ...string) error {
	if h.HasAnyFlag(flagKeys...) {
		return nil
	}
	return fmt.Errorf("at least one flag must be specified")
}

func (h *ValidateFlagAdapter) ValidateHasMoreArgsThan(count uint8) error {
	if h.HasMoreArgsThan(count) {
		return nil
	}
	return fmt.Errorf("at least %d arguments must be specified", count)
}

func (h *ValidateFlagAdapter) ValidateRequireFlag(flagKey string) error {
	if h.IsFlagChanged(flagKey) || !h.IsFlagRequired(flagKey) {
		return nil
	}
	return fmt.Errorf("flag --%s is required", flagKey)
}

func (h *ValidateFlagAdapter) ValidateStringFlag(flagKey string, fn func(string) error) error {
	if err := h.ValidateRequireFlag(flagKey); err != nil {
		return err
	}
	if !h.IsFlagChanged(flagKey) {
		return nil
	}

	if val, err := h.command.Flags().GetString(flagKey); err != nil {
		return fmt.Errorf("failed to get --%s flag value: %w", flagKey, err)
	} else if err = fn(val); err != nil {
		return fmt.Errorf("invalid value in --%s flag: %w", flagKey, err)
	}

	return nil
}

func (h *ValidateFlagAdapter) ValidateStringSliceFlag(flagKey string, fn func(string) error) []error {
	if err := h.ValidateRequireFlag(flagKey); err != nil {
		return []error{err}
	}
	if !h.IsFlagChanged(flagKey) {
		return nil
	}

	var errs []error
	if vals, err := h.command.Flags().GetStringSlice(flagKey); err != nil {
		errs = append(errs, fmt.Errorf("failed to get --%s flag value: %w", flagKey, err))
	} else {
		for i := range vals {
			if err = fn(vals[i]); err != nil {
				errs = append(errs, fmt.Errorf("invalid value in --%s flag: %w", flagKey, err))
			}
		}
	}

	return errs
}
