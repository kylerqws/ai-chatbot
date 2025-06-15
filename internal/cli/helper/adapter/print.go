package adapter

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// PrintAdapter provides the implementation for CLI adapter with output handling.
type PrintAdapter struct {
	command *cobra.Command
}

// NewPrintAdapter creates a new PrintAdapter adapter.
func NewPrintAdapter(cmd *cobra.Command) *PrintAdapter {
	return &PrintAdapter{command: cmd}
}

// PrintMessage outputs a general message to the default writer.
func (a *PrintAdapter) PrintMessage(args ...any) error {
	return a.PrintMessageToWriter(a.command.OutOrStdout(), args...)
}

// PrintErrMessage outputs an error message to the default writer.
func (a *PrintAdapter) PrintErrMessage(args ...any) error {
	return a.PrintMessageToWriter(a.command.ErrOrStderr(), args...)
}

// PrintMessageToWriter outputs a message to the specified writer.
func (*PrintAdapter) PrintMessageToWriter(w io.Writer, args ...any) error {
	if _, err := fmt.Fprintln(w, args...); err != nil {
		return err
	}
	return nil
}
