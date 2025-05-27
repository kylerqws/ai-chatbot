package openai

import (
	"github.com/spf13/cobra"

	//action "github.com/kylerqws/chatbot/cmd/openai/chat"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type ChatAdapter struct {
	*helper.ParentAdapter
}

func NewChatAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &ChatAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapter = helper.NewParentAdapter(app, cmd)
	return adp
}

func (a *ChatAdapter) Configure() *cobra.Command {
	//app := a.App()

	a.SetUse("chat")
	a.SetShort("Operations on chat management")

	a.AddChildren()

	return a.MainConfigure()
}
