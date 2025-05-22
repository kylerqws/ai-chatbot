package helper

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

type PrintAdapterHelper struct {
	command *cobra.Command
}

func NewPrintAdapterHelper(cmd *cobra.Command) *PrintAdapterHelper {
	return &PrintAdapterHelper{command: cmd}
}

func (h *PrintAdapterHelper) PrintMessage(args ...any) error {
	return h.PrintMessageToWriter(h.command.OutOrStdout(), args...)
}

func (h *PrintAdapterHelper) PrintErrMessage(args ...any) error {
	return h.PrintMessageToWriter(h.command.ErrOrStderr(), args...)
}

func (h *PrintAdapterHelper) PrintMessageToWriter(w io.Writer, args ...any) error {
	if _, err := fmt.Fprintln(w, args...); err != nil {
		return err
	}
	return nil
}
