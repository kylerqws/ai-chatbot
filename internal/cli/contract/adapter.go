package contract

import (
	"io"

	intapp "github.com/kylerqws/chatbot/internal/app"
	"github.com/spf13/cobra"
)

type Adapter interface {
	App() *intapp.App
	Command() *cobra.Command
	Configure() *cobra.Command
}

type ParentAdapter interface {
	Adapter
	Children() []*cobra.Command
}

type RootAdapter interface {
	ParentAdapter
	Version() string
}

type CommandAdapter interface {
	Adapter

	Errors() []error
	AddError(error)

	ShowErrors() bool
	ExistErrors() bool

	PrintErrors() error
	PrintErrorsToWriter(io.Writer) error
}
