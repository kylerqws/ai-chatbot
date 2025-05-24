package adapter

import (
	"fmt"
	"io"
	"strings"

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

func (h *ErrorAdapterHelper) ExistErrors() bool {
	return len(h.errors) > 0
}

func (h *ErrorAdapterHelper) AddError(err error) {
	if err != nil {
		h.errors = append(h.errors, err)
	}
}

func (h *ErrorAdapterHelper) AddErrors(errs ...error) {
	for i := range errs {
		if errs[i] != nil {
			h.errors = append(h.errors, errs[i])
		}
	}
}

func (h *ErrorAdapterHelper) ShowErrors() bool {
	show, err := h.command.Flags().GetBool("error")
	return show && err == nil
}

func (h *ErrorAdapterHelper) StringErrors() string {
	var buf strings.Builder
	if err := h.PrintErrorsToWriter(&buf); err != nil {
		return err.Error()
	}

	return buf.String()
}

func (h *ErrorAdapterHelper) ErrorIfExist(format string, args ...any) error {
	if !h.ExistErrors() {
		return nil
	}
	if !h.ShowErrors() {
		return fmt.Errorf(format, args...)
	}

	return fmt.Errorf("%s", h.StringErrors())
}

func (h *ErrorAdapterHelper) PrintErrors() error {
	return h.PrintErrorsToWriter(h.command.ErrOrStderr())
}

func (h *ErrorAdapterHelper) PrintErrorsToWriter(w io.Writer) error {
	count := len(h.errors) - 1

	for i := range h.errors {
		if i == 0 {
			if _, err := fmt.Fprintln(w); err != nil {
				return err
			}
		}

		msg := fmt.Errorf("- error: %w", h.errors[i])
		if _, err := fmt.Fprint(w, msg); err != nil {
			return err
		}

		if i < count {
			if _, err := fmt.Fprintln(w); err != nil {
				return err
			}
		}
	}

	return nil
}
