package adapter

import (
	"github.com/kylerqws/chatbot/internal/cli/contract"
	"io"
)

type CommandAdapter interface {
	contract.Adapter

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

	FuncArgs() contract.FuncArgs
	SetFuncArgs(contract.FuncArgs)

	FuncRunE() contract.FuncRunE
	SetFuncRunE(contract.FuncRunE)
}
