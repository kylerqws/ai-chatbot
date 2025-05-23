package openai

import (
	"github.com/spf13/cobra"

	//action "github.com/kylerqws/chatbot/cmd/openai/chat"
	intapp "github.com/kylerqws/chatbot/internal/app"
	hlppar "github.com/kylerqws/chatbot/internal/cli/helper/adapter/parent"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

type ChatAdapter struct {
	*hlppar.ParentAdapterHelper
}

func NewChatAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &ChatAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper =
		hlppar.NewParentAdapterHelper(app, cmd)

	return adp
}

func (a *ChatAdapter) Configure() *cobra.Command {
	//app := a.App()

	a.SetUse("chat")
	a.SetShort("Manage chats via the OpenAI API")

	a.AddChildren()

	return a.MainConfigure()
}
