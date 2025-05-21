package contract

import (
	"github.com/spf13/cobra"
	"io"
)

type CommandAdapter interface {
	Adapter

	Errors() []error
	AddError(error)

	ShowErrors() bool
	ExistErrors() bool

	PrintErrors() error
	PrintErrorsToWriter(io.Writer) error

	PrintMessage(...any) error
	PrintErrMessage(...any) error
	PrintMessageToWriter(io.Writer, ...any) error

	FuncArgs(*cobra.Command, []string) error
	SetFuncArgs(FuncArgs)

	FuncRunE(*cobra.Command, []string) error
	SetFuncRunE(FuncRunE)
}
