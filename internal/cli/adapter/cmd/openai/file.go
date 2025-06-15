package openai

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/openai/file"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// FileAdapter provides the implementation for the file command adapter.
type FileAdapter struct {
	*helper.ParentAdapter
}

// NewFileAdapter creates a new file command adapter.
func NewFileAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &FileAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapter = helper.NewParentAdapter(app, cmd)
	return adp
}

// Configure applies configuration for the command.
func (a *FileAdapter) Configure() *cobra.Command {
	app := a.App()

	a.SetUse("file")
	a.SetShort("Operations on files management")

	a.AddChildren(
		action.ListCommand(app),
		action.UploadCommand(app),
		action.DeleteCommand(app),
	)

	return a.MainConfigure()
}
