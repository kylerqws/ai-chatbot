package helper

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

type ErrorAdapterHelper struct {
	command *cobra.Command
	errors  []error
}

func NewErrorAdapterHelper(cmd *cobra.Command) *ErrorAdapterHelper {
	return &ErrorAdapterHelper{command: cmd}
}

func (h *ErrorAdapterHelper) Errors() []error {
	return h.errors
}

func (h *ErrorAdapterHelper) AddErrors(errs ...error) {
	h.errors = append(h.errors, errs...)
}

func (h *ErrorAdapterHelper) AddError(err error) {
	h.AddErrors(err)
}

func (h *ErrorAdapterHelper) ShowErrors() bool {
	show, err := h.command.Flags().GetBool("error")
	return show && err == nil
}

func (h *ErrorAdapterHelper) ExistErrors() bool {
	return len(h.errors) > 0
}

func (h *ErrorAdapterHelper) PrintErrors() error {
	return h.PrintErrorsToWriter(h.command.ErrOrStderr())
}

func (h *ErrorAdapterHelper) PrintErrorsToWriter(w io.Writer) error {
	pfx := ""
	if len(h.errors) > 1 {
		pfx = "- "
	}

	for i := range h.errors {
		msg := fmt.Errorf("%serror: %w\n", pfx, h.errors[i])
		if _, err := fmt.Fprint(w, msg); err != nil {
			return err
		}
	}

	return nil
}
