package openai

import (
	"github.com/spf13/cobra"

	// TODO: will use it: action "github.com/kylerqws/chatbot/cmd/openai/chat"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// ChatAdapter provides the implementation for the OpenAI chat CLI adapter.
type ChatAdapter struct {
	*helper.ParentAdapter
}

// NewChatAdapter creates a new ChatAdapter adapter.
func NewChatAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &ChatAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapter = helper.NewParentAdapter(app, cmd)
	return adp
}

// Configure applies configuration for the command.
func (a *ChatAdapter) Configure() *cobra.Command {
	// TODO: will use it: app := a.App()

	a.SetUse("chat")
	a.SetShort("Chat completion via OpenAI API")

	a.AddChildren( /* TODO: add subcommands*/ )

	return a.MainConfigure()
}
