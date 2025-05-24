package adapter

import (
	"github.com/kylerqws/chatbot/internal/app/helper"
	"github.com/spf13/cobra"
)

const (
	ExecStatusSuccess = "✓"
	ExecStatusFailed  = "✗"
)

type FormatAdapterHelper struct {
	command *cobra.Command
}

func NewFormatAdapterHelper(cmd *cobra.Command) *FormatAdapterHelper {
	return &FormatAdapterHelper{command: cmd}
}

func (*FormatAdapterHelper) FormatBytes(val int64, empty *string) string {
	if val == 0 && empty != nil {
		return *empty
	}
	return helper.FormatBytes(val)
}

func (*FormatAdapterHelper) FormatTime(val int64, empty *string) string {
	if val == 0 && empty != nil {
		return *empty
	}
	return helper.FormatTime(val)
}

func (*FormatAdapterHelper) FormatExecStatus(status bool) string {
	if status {
		return ExecStatusSuccess
	}
	return ExecStatusFailed
}

func (*FormatAdapterHelper) FormatString(val string, empty *string) string {
	if val == "" && empty != nil {
		return *empty
	}
	return val
}
