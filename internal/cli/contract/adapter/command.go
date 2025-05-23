package adapter

import (
	"io"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type CommandAdapter interface {
	ctr.Adapter

	Errors() []error
	AddErrors(...error)
	AddError(error)

	ShowErrors() bool
	ExistErrors() bool

	PrintErrors() error
	PrintErrorsToWriter(io.Writer) error

	PrintMessage(...any) error
	PrintErrMessage(...any) error
	PrintMessageToWriter(io.Writer, ...any) error

	FuncArgs() ctr.FuncArgs
	SetFuncArgs(ctr.FuncArgs)

	FuncRunE() ctr.FuncRunE
	SetFuncRunE(ctr.FuncRunE)
}
