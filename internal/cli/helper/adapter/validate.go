package adapter

import (
	"fmt"
	"log"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
	"github.com/spf13/cobra"
)

type ValidateAdapter struct {
	command *cobra.Command
}

func NewValidateAdapter(cmd *cobra.Command) *ValidateAdapter {
	return &ValidateAdapter{command: cmd}
}

func (h *ValidateAdapter) HasChangedAnyFlag(names ...string) bool {
	for i := range names {
		if h.IsFlagChanged(names[i]) {
			return true
		}
	}
	return false
}

func (h *ValidateAdapter) HasMoreArgsThan(count uint8) bool {
	return len(h.command.Flags().Args()) > int(count)
}

func (h *ValidateAdapter) IsFlagChanged(key string) bool {
	return h.command.Flags().Changed(key)
}

func (h *ValidateAdapter) IsFlagRequired(key string) bool {
	f := h.command.Flags().Lookup(key)
	return f != nil && f.Annotations != nil && f.Annotations[cobra.BashCompOneRequiredFlag] != nil
}

func (h *ValidateAdapter) ValidateRequireFlag(key string) error {
	if h.IsFlagChanged(key) || !h.IsFlagRequired(key) {
		return nil
	}
	return h.errorRequiredFlag(key)
}

func (h *ValidateAdapter) ValidateHasChangedAnyFlag(keys ...string) error {
	if h.HasChangedAnyFlag(keys...) {
		return nil
	}
	return h.errorMustBeSpecifiedOneFlag()
}

func (h *ValidateAdapter) ValidateHasMoreArgsThan(count uint8) error {
	if h.HasMoreArgsThan(count) {
		return nil
	}
	return h.errorMustBeSpecifiedArgsThan(count)
}

func (h *ValidateAdapter) ValidateStringFlag(key string, fn ctr.FuncValidateString) error {
	if err := h.ValidateRequireFlag(key); err != nil {
		return err
	}
	if !h.IsFlagChanged(key) {
		return nil
	}

	val, err := h.command.Flags().GetString(key)
	if err != nil {
		return h.errorFailedToGetFlag(key, err)
	}
	if err = fn(val); err != nil {
		return h.errorInvalidValueFlag(key, err)
	}

	return nil
}

func (h *ValidateAdapter) ValidateStringSliceFlag(key string, fn ctr.FuncValidateString) []error {
	if err := h.ValidateRequireFlag(key); err != nil {
		return []error{err}
	}
	if !h.IsFlagChanged(key) {
		return nil
	}

	vals, err := h.command.Flags().GetStringSlice(key)
	if err != nil {
		return []error{h.errorFailedToGetFlag(key, err)}
	}

	var errs []error
	for i := range vals {
		if err = fn(vals[i]); err != nil {
			errs = append(errs, h.errorInvalidValueFlag(key, err))
		}
	}

	return errs
}

func (h *ValidateAdapter) ValidateUint8Flag(key string, fn ctr.FuncValidateUint8) error {
	if err := h.ValidateRequireFlag(key); err != nil {
		return err
	}
	if !h.IsFlagChanged(key) {
		return nil
	}

	val, err := h.command.Flags().GetUint8(key)
	if err != nil {
		return h.errorFailedToGetFlag(key, err)
	}
	if err = fn(val); err != nil {
		return h.errorInvalidValueFlag(key, err)
	}

	return nil
}

func (*ValidateAdapter) ValidateUint8SliceFlag(_ string, _ ctr.FuncValidateUint8) []error {
	// TODO: need to implement GetUint8Slice as it is not implemented in the Cobra package
	log.Fatalf("validation uint8 slice flag is not implemented")

	return []error{}
}

func (h *ValidateAdapter) ValidateUintFlag(key string, fn ctr.FuncValidateUint) error {
	if err := h.ValidateRequireFlag(key); err != nil {
		return err
	}
	if !h.IsFlagChanged(key) {
		return nil
	}

	val, err := h.command.Flags().GetUint(key)
	if err != nil {
		return h.errorFailedToGetFlag(key, err)
	}
	if err = fn(val); err != nil {
		return h.errorInvalidValueFlag(key, err)
	}

	return nil
}

func (h *ValidateAdapter) ValidateUintSliceFlag(key string, fn ctr.FuncValidateUint) []error {
	if err := h.ValidateRequireFlag(key); err != nil {
		return []error{err}
	}
	if !h.IsFlagChanged(key) {
		return nil
	}

	vals, err := h.command.Flags().GetUintSlice(key)
	if err != nil {
		return []error{h.errorFailedToGetFlag(key, err)}
	}

	var errs []error
	for i := range vals {
		if err = fn(vals[i]); err != nil {
			errs = append(errs, h.errorInvalidValueFlag(key, err))
		}
	}

	return errs
}

func (*ValidateAdapter) errorRequiredFlag(key string) error {
	return fmt.Errorf("flag --%s is required", key)
}

func (*ValidateAdapter) errorMustBeSpecifiedOneFlag() error {
	return fmt.Errorf("at least one flag must be specified")
}

func (*ValidateAdapter) errorMustBeSpecifiedArgsThan(count uint8) error {
	return fmt.Errorf("at least %d argument(s) must be specified", count)
}

func (*ValidateAdapter) errorFailedToGetFlag(key string, err error) error {
	return fmt.Errorf("failed to get --%s flag: %w", key, err)
}

func (*ValidateAdapter) errorInvalidValueFlag(key string, err error) error {
	return fmt.Errorf("invalid value in --%s flag: %w", key, err)
}
