package openai

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type JobAdapter struct {
	app      *app.App
	command  *cobra.Command
	children []*cobra.Command
}

func NewJobAdapter(app *app.App) ctr.ParentAdapter {
	return &JobAdapter{app: app}
}

func (a *JobAdapter) Configure() *cobra.Command {
	a.command = &cobra.Command{
		Use:   "job",
		Short: "Manage jobs used with the OpenAI API",
	}

	a.children = []*cobra.Command{}

	return setup.ParentConfigure(a)
}

func (a *JobAdapter) App() *app.App {
	return a.app
}

func (a *JobAdapter) Command() *cobra.Command {
	return a.command
}

func (a *JobAdapter) Children() []*cobra.Command {
	return a.children
}
