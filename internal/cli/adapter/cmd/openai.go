package cmd

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/openai"
	intapp "github.com/kylerqws/chatbot/internal/app"
	inthlp "github.com/kylerqws/chatbot/internal/cli/helper"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type OpenAIAdapter struct {
	*inthlp.ParentAdapterHelper
}

func NewOpenAIAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &OpenAIAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper =
		inthlp.NewParentAdapterHelper(adp, app, cmd)

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
