package contract

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
)

type Adapter interface {
	App() *app.App
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
}
