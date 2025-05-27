package openai

import (
	"github.com/spf13/cobra"

	//action "github.com/kylerqws/chatbot/cmd/openai/job"
	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

type JobAdapter struct {
	*helper.ParentAdapter
}

func NewJobAdapter(app *intapp.App) ctr.ParentAdapter {
	adp := &JobAdapter{}
	cmd := &cobra.Command{}

	adp.ParentAdapter = helper.NewParentAdapter(app, cmd)
	return adp
}

func (a *JobAdapter) Configure() *cobra.Command {
	//app := a.App()

	a.SetUse("job")
	a.SetShort("Operations on job management")

	a.AddChildren()

	return a.MainConfigure()
}
