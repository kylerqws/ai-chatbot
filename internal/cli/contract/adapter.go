package contract

import (
	"github.com/kylerqws/chatbot/internal/app"
	"github.com/spf13/cobra"
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
}

type CommandAdapter interface {
	Adapter
	Errors() []error
}
