package openai

import (
	"github.com/spf13/cobra"

	//action "github.com/kylerqws/chatbot/cmd/openai/job"
	intapp "github.com/kylerqws/chatbot/internal/app"
	inthlp "github.com/kylerqws/chatbot/internal/cli/helper"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type JobAdapter struct {
	*inthlp.ParentAdapterHelper
}

func NewJobAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &JobAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapterHelper =
		inthlp.NewParentAdapterHelper(adp, app, cmd)

	return adp
}

func (a *JobAdapter) Configure() *cobra.Command {
	_ = a.App()

	a.SetUse("job")
	a.SetShort("Manage jobs used with the OpenAI API")
	a.AddChildren()

	return a.MainConfigure()
}
