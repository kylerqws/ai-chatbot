package openai

import (
	"github.com/spf13/cobra"

	action "github.com/kylerqws/chatbot/cmd/openai/model"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// ModelAdapter provides the implementation for the model CLI adapter.
type ModelAdapter struct {
	*helper.ParentAdapter
}

// NewModelAdapter creates a new model command adapter.
func NewModelAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &ModelAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapter = helper.NewParentAdapter(app, cmd)
	return adp
}

// Configure applies configuration for the command.
func (a *ModelAdapter) Configure() *cobra.Command {
	app := a.App()

	a.SetUse("model")
	a.SetShort("Operations on models management")

	a.AddChildren(
		action.ListCommand(app),
		action.DeleteCommand(app),
	)

	return a.MainConfigure()
}
