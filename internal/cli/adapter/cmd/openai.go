package cmd

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/openai"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// OpenAIAdapter provides the implementation for the OpenAI API CLI adapter.
type OpenAIAdapter struct {
	*helper.ParentAdapter
}

// NewOpenAIAdapter creates a new OpenAIAdapter adapter.
func NewOpenAIAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &OpenAIAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapter = helper.NewParentAdapter(app, cmd)
	return adp
}

// Configure applies configuration for the command.
func (a *OpenAIAdapter) Configure() *cobra.Command {
	app := a.App()

	a.SetUse("openai")
	a.SetShort("OpenAI API integration commands")

	a.AddChildren(
		action.FileCommand(app),
		action.FineTuningCommand(app),
		action.ModelCommand(app),
		action.ChatCommand(app),
	)

	return a.MainConfigure()
}
