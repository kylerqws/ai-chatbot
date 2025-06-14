package adapter

import (
	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
	"io"
)

// CommandAdapter defines the interface for a CLI adapters with logic and output.
type CommandAdapter interface {
	GeneralAdapter

	// Errors returns the list of collected errors.
	Errors() []error

	// ExistErrors reports whether any errors have been recorded.
	ExistErrors() bool

	// AddError adds a single error to the collection.
	AddError(err error)

	// AddErrors adds multiple errors to the collection.
	AddErrors(errs ...error)

	// ShowErrors reports whether errors should be displayed to the user.
	ShowErrors() bool

	// StringErrors returns all errors as a single concatenated string.
	StringErrors() string

	// ErrorIfExist returns a formatted error if any exist.
	ErrorIfExist(format string, args ...any) error

	// PrintErrors outputs all errors to the default writer.
	PrintErrors() error

	// PrintErrorsToWriter outputs all errors to the specified writer.
	PrintErrorsToWriter(w io.Writer) error

	// PrintMessage outputs a general message to the default writer.
	PrintMessage(args ...any) error

	// PrintErrMessage outputs an error message to the default writer.
	PrintErrMessage(args ...any) error

	// PrintMessageToWriter outputs a message to the specified writer.
	PrintMessageToWriter(w io.Writer, args ...any) error

	// FuncArgs returns the cobra-compatible command argument handler.
	FuncArgs() ctr.FuncArgs

	// SetFuncArgs sets the cobra-compatible command argument handler.
	SetFuncArgs(handler ctr.FuncArgs)

	// FuncRunE returns the cobra-compatible command execution function.
	FuncRunE() ctr.FuncRunE

	// SetFuncRunE sets the cobra-compatible command execution function.
	SetFuncRunE(handler ctr.FuncRunE)
}
