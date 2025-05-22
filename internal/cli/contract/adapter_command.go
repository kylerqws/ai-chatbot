package contract

import "io"

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

	FuncArgs() FuncArgs
	SetFuncArgs(FuncArgs)

	FuncRunE() FuncRunE
	SetFuncRunE(FuncRunE)
}
