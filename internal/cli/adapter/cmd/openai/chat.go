package openai

import (
	"github.com/spf13/cobra"

	//action "github.com/kylerqws/chatbot/cmd/openai/chat"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type ChatAdapter struct {
	*helper.ParentAdapterHelper
}

func NewChatAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &ChatAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper = helper.NewParentAdapterHelper(app, cmd)
	return adp
}

func (a *ChatAdapter) Configure() *cobra.Command {
	//app := a.App()

	a.SetUse("chat")
	a.SetShort("Manage chats via the OpenAI API")

	a.AddChildren()

	return a.MainConfigure()
}
