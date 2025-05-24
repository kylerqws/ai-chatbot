package contract

import (
	"io"

	intapp "github.com/kylerqws/chatbot/internal/app"
	"github.com/spf13/cobra"
)

type GeneralAdapter interface {
	App() *intapp.App
	Command() *cobra.Command

	Use() string
	SetUse(string)

	Short() string
	SetShort(string)

	Configure() *cobra.Command
	MainConfigure() *cobra.Command
}

type ParentAdapter interface {
	GeneralAdapter

	Children() []*cobra.Command
	ExistChildren() bool

	AddChildren(...*cobra.Command)
	AddChild(*cobra.Command)
}

type RootAdapter interface {
	ParentAdapter

	Version() string
	SetVersion(string)
}

type CommandAdapter interface {
	GeneralAdapter

	Errors() []error
	AddErrors(...error)
	AddError(error)

	ShowErrors() bool
	ExistErrors() bool
	ErrorIfExist(string, ...any) error

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
