package adapter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// ErrorAdapter provides the implementation for CLI error handling.
type ErrorAdapter struct {
	command *cobra.Command
	errors  []error
}

// NewErrorAdapter creates a new error adapter.
func NewErrorAdapter(cmd *cobra.Command) *ErrorAdapter {
	return &ErrorAdapter{command: cmd}
}

// Errors returns the list of collected errors.
func (a *ErrorAdapter) Errors() []error {
	return a.errors
}

// ExistErrors reports whether any errors have been recorded.
func (a *ErrorAdapter) ExistErrors() bool {
	return len(a.errors) > 0
}

// AddError adds a single error to the collection.
func (a *ErrorAdapter) AddError(err error) {
	if err != nil {
		a.errors = append(a.errors, err)
	}
}

// AddErrors adds multiple errors to the collection.
func (a *ErrorAdapter) AddErrors(errs ...error) {
	for i := range errs {
		a.AddError(errs[i])
	}
}

// ShowErrors reports whether errors should be displayed to the user.
func (a *ErrorAdapter) ShowErrors() bool {
	show, err := a.command.Flags().GetBool("error")
	return show && err == nil
}

// StringErrors returns all errors as a single concatenated string.
func (a *ErrorAdapter) StringErrors() string {
	var buf strings.Builder
	if err := a.PrintErrorsToWriter(&buf); err != nil {
		return err.Error()
	}
	return buf.String()
}

// ErrorIfExist returns a formatted error if any exist.
func (a *ErrorAdapter) ErrorIfExist(format string, args ...any) error {
	if a.ExistErrors() {
		if !a.ShowErrors() {
			return fmt.Errorf(format, args...)
		}

		exec := fmt.Sprintf("%s %s", filepath.Base(os.Args[0]), strings.Join(os.Args[1:], " "))
		return fmt.Errorf("Failed to execute command: `%s`\n%s", exec, a.StringErrors())
	}
	return nil
}

// PrintErrors outputs all errors to the default output stream.
func (a *ErrorAdapter) PrintErrors() error {
	return a.PrintErrorsToWriter(a.command.ErrOrStderr())
}

// PrintErrorsToWriter outputs all errors to the specified writer.
func (a *ErrorAdapter) PrintErrorsToWriter(w io.Writer) error {
	count := len(a.errors) - 1
	for i := range a.errors {
		msg := fmt.Errorf("- error: %w", a.errors[i])
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
