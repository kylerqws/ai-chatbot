package helper

import (
	"github.com/spf13/cobra"

	hlpapp "github.com/kylerqws/chatbot/internal/app/helper"
)

const (
	execStatusSuccess = "✓"
	execStatusFailed  = "✗"
)

type FormatAdapterHelper struct {
	command *cobra.Command
}

func NewFormatAdapterHelper(cmd *cobra.Command) *FormatAdapterHelper {
	return &FormatAdapterHelper{command: cmd}
}

func (*FormatAdapterHelper) FormatBytes(b int64, empty *string) string {
	if b == 0 && empty != nil {
		return *empty
	}
	return hlpapp.FormatBytes(b)
}

func (*FormatAdapterHelper) FormatTime(t int64, empty *string) string {
	if t == 0 && empty != nil {
		return *empty
	}
	return hlpapp.FormatTime(t)
}

func (*FormatAdapterHelper) FormatExecStatus(v bool) string {
	if v {
		return execStatusSuccess
	}
	return execStatusFailed
}

func (*FormatAdapterHelper) FormatString(s string, empty *string) string {
	if s == "" && empty != nil {
		return *empty
	}
	return s
}
