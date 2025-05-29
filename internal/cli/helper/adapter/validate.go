package adapter

import (
	"fmt"
	"github.com/spf13/cobra"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type ValidateAdapter struct {
	command *cobra.Command
}

func NewValidateAdapter(cmd *cobra.Command) *ValidateAdapter {
	return &ValidateAdapter{command: cmd}
}

func (h *ValidateAdapter) HasAnyFlag(names ...string) bool {
	fgs := h.command.Flags()
	for i := range names {
		if fgs.Changed(names[i]) {
			return true
		}
	}
	return false
}

func (h *ValidateAdapter) HasMoreArgsThan(count uint8) bool {
	return len(h.command.Flags().Args()) > int(count)
}

func (h *ValidateAdapter) IsFlagChanged(flagKey string) bool {
	return h.command.Flags().Changed(flagKey)
}

func (h *ValidateAdapter) IsFlagRequired(flagKey string) bool {
	f := h.command.Flags().Lookup(flagKey)
	return f != nil && f.Annotations != nil && f.Annotations[cobra.BashCompOneRequiredFlag] != nil
}

func (h *ValidateAdapter) ValidateRequireFlag(flagKey string) error {
	if h.IsFlagChanged(flagKey) || !h.IsFlagRequired(flagKey) {
		return nil
	}
	return fmt.Errorf("flag --%s is required", flagKey)
}

func (h *ValidateAdapter) ValidateHasAnyFlags(flagKeys ...string) error {
	if h.HasAnyFlag(flagKeys...) {
		return nil
	}
	return fmt.Errorf("at least one flag must be specified")
}

func (h *ValidateAdapter) ValidateHasMoreArgsThan(count uint8) error {
	if h.HasMoreArgsThan(count) {
		return nil
	}
	return fmt.Errorf("at least %d arguments must be specified", count+1)
}

func (h *ValidateAdapter) ValidateStringFlag(flagKey string, fn ctr.FuncValidateString) error {
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

func (h *ValidateAdapter) ValidateStringSliceFlag(flagKey string, fn ctr.FuncValidateString) []error {
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
