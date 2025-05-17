package openai

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type ChatAdapter struct {
	app      *app.App
	command  *cobra.Command
	children []*cobra.Command
}

func NewChatAdapter(app *app.App) ctr.ParentAdapter {
	return &ChatAdapter{app: app}
}

func (a *ChatAdapter) Configure() *cobra.Command {
	a.command = &cobra.Command{
		Use:   "chat",
		Short: "Manage chats used with the OpenAI API",
	}

	a.children = []*cobra.Command{}

	return setup.ParentConfigure(a)
}

func (a *ChatAdapter) App() *app.App {
	return a.app
}

func (a *ChatAdapter) Command() *cobra.Command {
	return a.command
}

func (a *ChatAdapter) Children() []*cobra.Command {
	return a.children
}
