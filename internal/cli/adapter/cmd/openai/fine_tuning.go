package openai

import (
	"github.com/spf13/cobra"

	// TODO: will use it: action "github.com/kylerqws/chatbot/cmd/openai/fine-tuning"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// FineTuningAdapter provides the implementation for the OpenAI fine-tuning CLI adapter.
type FineTuningAdapter struct {
	*helper.ParentAdapter
}

// NewFineTuningAdapter creates a new FineTuningAdapter adapter.
func NewFineTuningAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &FineTuningAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapter = helper.NewParentAdapter(app, cmd)
	return adp
}

// Configure applies configuration for the command.
func (a *FineTuningAdapter) Configure() *cobra.Command {
	// TODO: will use it: app := a.App()

	a.SetUse("fine-tuning")
	a.SetShort("Operations on fine-tuning jobs management")

	a.AddChildren( /* TODO: add subcommands*/ )

	return a.MainConfigure()
}
