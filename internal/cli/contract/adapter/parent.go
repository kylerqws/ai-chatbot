package adapter

import (
	"github.com/spf13/cobra"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type ParentAdapter interface {
	ctr.Adapter

	Children() []*cobra.Command
	ExistChildren() bool

	AddChildren(...*cobra.Command)
	AddChild(*cobra.Command)
}
