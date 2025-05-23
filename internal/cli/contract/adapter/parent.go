package adapter

import (
	"github.com/kylerqws/chatbot/internal/cli/contract"
	"github.com/spf13/cobra"
)

type ParentAdapter interface {
	contract.Adapter

	Children() []*cobra.Command
	ExistChildren() bool

	AddChildren(...*cobra.Command)
	AddChild(*cobra.Command)
}
