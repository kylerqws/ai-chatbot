package openai

import (
	"github.com/spf13/cobra"

	//action "github.com/kylerqws/chatbot/cmd/openai/chat"
	intapp "github.com/kylerqws/chatbot/internal/app"
	inthlp "github.com/kylerqws/chatbot/internal/cli/helper"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type ChatAdapter struct {
	*inthlp.ParentAdapterHelper
}

func NewChatAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &ChatAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper =
		inthlp.NewParentAdapterHelper(adp, app, cmd)

	return adp
}

func (a *ChatAdapter) Configure() *cobra.Command {
	_ = a.App()

	a.SetUse("chat")
	a.SetShort("Manage chats used with the OpenAI API")
	a.AddChildren()

	return a.MainConfigure()
}
