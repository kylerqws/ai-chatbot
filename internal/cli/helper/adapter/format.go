package adapter

import (
	"github.com/kylerqws/chatbot/internal/app/helper"
	"github.com/spf13/cobra"
)

const (
	ExecStatusSuccess = "✓"
	ExecStatusFailed  = "✗"
)

type FormatAdapter struct {
	command *cobra.Command
}

func NewFormatAdapter(cmd *cobra.Command) *FormatAdapter {
	return &FormatAdapter{command: cmd}
}

func (*FormatAdapter) FormatBytes(val int64, empty *string) string {
	if val == 0 && empty != nil {
		return *empty
	}
	return helper.FormatBytes(val)
}

func (*FormatAdapter) FormatTime(val int64, empty *string) string {
	if val == 0 && empty != nil {
		return *empty
	}
	return helper.FormatTime(val)
}

func (*FormatAdapter) FormatExecStatus(status bool) string {
	if status {
		return ExecStatusSuccess
	}
	return ExecStatusFailed
}

func (*FormatAdapter) FormatString(val string, empty *string) string {
	if val == "" && empty != nil {
		return *empty
	}
	return val
}
