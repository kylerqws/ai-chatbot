package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	"github.com/kylerqws/chatbot/cmd/openai"
	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type OpenAIAdapter struct {
	app      *app.App
	command  *cobra.Command
	children []*cobra.Command
}

func NewOpenAIAdapter(app *app.App) ctr.ParentAdapter {
	return &OpenAIAdapter{app: app}
}

func (a *OpenAIAdapter) Configure() *cobra.Command {
	a.command = &cobra.Command{
		Use:   "openai",
		Short: "Commands for interacting with the OpenAI API",
	}

	a.children = []*cobra.Command{
		openai.FileCommand(a.app),
		openai.JobCommand(a.app),
		openai.ChatCommand(a.app),
	}

	return setup.ParentConfigure(a)
}

func (a *OpenAIAdapter) App() *app.App {
	return a.app
}

func (a *OpenAIAdapter) Command() *cobra.Command {
	return a.command
}

func (a *OpenAIAdapter) Children() []*cobra.Command {
	return a.children
}
