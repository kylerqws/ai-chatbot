package cmd

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/openai"
	intapp "github.com/kylerqws/chatbot/internal/app"
	hlppar "github.com/kylerqws/chatbot/internal/cli/helper/adapter/parent"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

type OpenAIAdapter struct {
	*hlppar.ParentAdapterHelper
}

func NewOpenAIAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &OpenAIAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper =
		hlppar.NewParentAdapterHelper(app, cmd)

	return adp
}

func (a *OpenAIAdapter) Configure() *cobra.Command {
	app := a.App()

	a.SetUse("openai")
	a.SetShort("OpenAI API integration commands")

	a.AddChildren(
		action.FileCommand(app),
		action.JobCommand(app),
		action.ChatCommand(app),
	)

	return a.MainConfigure()
}
