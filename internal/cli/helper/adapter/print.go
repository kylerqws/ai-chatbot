package adapter

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

type PrintAdapter struct {
	command *cobra.Command
}

func NewPrintAdapter(cmd *cobra.Command) *PrintAdapter {
	return &PrintAdapter{command: cmd}
}

func (h *PrintAdapter) PrintMessage(args ...any) error {
	return h.PrintMessageToWriter(h.command.OutOrStdout(), args...)
}

func (h *PrintAdapter) PrintErrMessage(args ...any) error {
	return h.PrintMessageToWriter(h.command.ErrOrStderr(), args...)
}

func (*PrintAdapter) PrintMessageToWriter(w io.Writer, args ...any) error {
	if _, err := fmt.Fprintln(w, args...); err != nil {
		return err
	}
	return nil
}
