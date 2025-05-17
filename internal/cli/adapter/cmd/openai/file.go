package openai

import (
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/cli/setup"

	"github.com/kylerqws/chatbot/cmd/openai/file"
	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type FileAdapter struct {
	app      *app.App
	command  *cobra.Command
	children []*cobra.Command
}

func NewFileAdapter(app *app.App) ctr.ParentAdapter {
	return &FileAdapter{app: app}
}

func (a *FileAdapter) Configure() *cobra.Command {
	a.command = &cobra.Command{
		Use:   "file",
		Short: "Manage files used with the OpenAI API",
	}

	a.children = []*cobra.Command{
		file.ListCommand(a.app),
	}

	return setup.ParentConfigure(a)
}

func (a *FileAdapter) App() *app.App {
	return a.app
}

func (a *FileAdapter) Command() *cobra.Command {
	return a.command
}

func (a *FileAdapter) Children() []*cobra.Command {
	return a.children
}
