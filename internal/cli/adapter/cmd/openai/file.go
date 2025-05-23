package openai

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/openai/file"
	intapp "github.com/kylerqws/chatbot/internal/app"
	hlppar "github.com/kylerqws/chatbot/internal/cli/helper/adapter/parent"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

type FileAdapter struct {
	*hlppar.ParentAdapterHelper
}

func NewFileAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &FileAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper =
		hlppar.NewParentAdapterHelper(app, cmd)

	return adp
}

func (a *FileAdapter) Configure() *cobra.Command {
	app := a.App()

	a.SetUse("file")
	a.SetShort("Manage files via the OpenAI API")

	a.AddChildren(
		action.ListCommand(app),
		action.UploadCommand(app),
		action.DeleteCommand(app),
	)

	return a.MainConfigure()
}
