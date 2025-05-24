package openai

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/openai/file"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type FileAdapter struct {
	*helper.ParentAdapterHelper
}

func NewFileAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &FileAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper = helper.NewParentAdapterHelper(app, cmd)
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
