package adapter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type ErrorAdapter struct {
	command *cobra.Command
	errors  []error
}

func NewErrorAdapter(cmd *cobra.Command) *ErrorAdapter {
	return &ErrorAdapter{command: cmd}
}

func (h *ErrorAdapter) Errors() []error {
	return h.errors
}

func (h *ErrorAdapter) ExistErrors() bool {
	return len(h.errors) > 0
}

func (h *ErrorAdapter) AddError(err error) {
	if err != nil {
		h.errors = append(h.errors, err)
	}
}

func (h *ErrorAdapter) AddErrors(errs ...error) {
	for i := range errs {
		h.AddError(errs[i])
	}
}

func (h *ErrorAdapter) ShowErrors() bool {
	show, err := h.command.Flags().GetBool("error")
	return show && err == nil
}

func (h *ErrorAdapter) StringErrors() string {
	var buf strings.Builder
	if err := h.PrintErrorsToWriter(&buf); err != nil {
		return err.Error()
	}

	return buf.String()
}

func (h *ErrorAdapter) ErrorIfExist(format string, args ...any) error {
	if h.ExistErrors() {
		if !h.ShowErrors() {
			return fmt.Errorf(format, args...)
		}

		exec := fmt.Sprintf("%s %s", filepath.Base(os.Args[0]), strings.Join(os.Args[1:], " "))
		return fmt.Errorf("Failed to execute command: `%s`\n%s", exec, h.StringErrors())
	}

	return nil
}

func (h *ErrorAdapter) PrintErrors() error {
	return h.PrintErrorsToWriter(h.command.ErrOrStderr())
}

func (h *ErrorAdapter) PrintErrorsToWriter(w io.Writer) error {
	count := len(h.errors) - 1

	for i := range h.errors {
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
