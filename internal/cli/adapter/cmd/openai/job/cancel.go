package job

import (
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctradp "github.com/kylerqws/chatbot/internal/cli/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type CancelAdapter struct {
	*ListAdapter
}

func NewCancelAdapter(app *intapp.App) ctradp.CommandAdapter {
	adp := &CancelAdapter{}
	adp.ListAdapter = NewListAdapter(app).(*ListAdapter)

	return adp
}

func (a *CancelAdapter) Configure() *cobra.Command {
	a.SetUse("cancel <filter-flag> [filter-flag...]")
	a.SetShort("Cancel one or more jobs in OpenAI account")
	a.SetLong(a.JobCancelHelpInfo(a.OpenAiAdapter))

	a.SetFuncArgs(a.Validate)
	a.SetFuncRunE(a.Cancel)

	a.ConfigureFilters()
	a.ConfigureFlags()

	return a.MainConfigure()
}

func (a *CancelAdapter) ConfigureFilters() {
	a.AddFilterKey(helper.AllFlagKey)
	a.ListAdapter.ConfigureFilters()
}

func (a *CancelAdapter) ConfigureFlags() {
	desc := "Cancel all jobs (overrides other filters)\n"
	a.AddBoolFlag(helper.AllFlagKey, "", false, desc)

	a.ListAdapter.ConfigureFlags()
}

func (a *CancelAdapter) Validate(cmd *cobra.Command, args []string) error {
	a.AddErrors(a.ValidateHasChangedAnyFlag(a.FilterKeys()...))
	if err := a.ListAdapter.Validate(cmd, args); err != nil {
		return err
	}

	return a.ErrorIfExist("One or more arguments/flags are invalid or missing.")
}

func (a *CancelAdapter) Cancel(_ *cobra.Command, _ []string) error {
	if a.Request() && a.ExistJobs() {
		if err := a.PrintJobs(); err != nil {
			return err
		}
	} else if !a.ExistErrors() {
		return a.PrintMessage("No jobs found.")
	}

	return a.ErrorIfExist("Failed to cancel jobs or data is unavailable.")
}

func (a *CancelAdapter) Request() bool {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().JobService()

	if !a.ListAdapter.Request() {
		return false
	}
	jobs := a.Jobs()

	for i := range jobs {
		resp, err := svc.CancelJob(ctx, &ctrsvc.CancelJobRequest{
			JobID: jobs[i].ID,
		})

		if err != nil {
			a.AddError(err)
		}
		stsCodes := a.JobStatusManager().Codes

		jobs[i].ExecStatus =
			resp.Job.Status == stsCodes.Cancelled || jobs[i].Status == stsCodes.Succeeded ||
				jobs[i].Status == stsCodes.Failed || jobs[i].Status == stsCodes.Cancelled

	}

	return true
}

func (a *CancelAdapter) PrintJobs() error {
	_ = a.CreateTable()

	a.AppendTableHeader("Job ID", "Training File", "Validation File",
		"Model", "Fine-Tuned Model", "Created", "Finished", "Status", "State")

	a.SetColumnTableConfigs(
		a.ColumnConfig(1, text.AlignCenter, 27, text.Colors{text.Bold}),
		a.ColumnConfig(2, text.AlignRight, 19, nil),
		a.ColumnConfig(3, text.AlignRight, 19, nil),
		a.ColumnConfig(4, text.AlignRight, 19, nil),
		a.ColumnConfig(5, text.AlignRight, 19, nil),
		a.ColumnConfig(6, text.AlignRight, 19, nil),
		a.ColumnConfig(7, text.AlignRight, 19, nil),
		a.ColumnConfig(8, text.AlignRight, 10, nil),
		a.ColumnConfig(9, text.AlignCenter, 7, text.Colors{text.Bold}),
	)

	jobs := a.Jobs()
	empty := helper.EmptyTableColumn

	for i := range jobs {
		a.AppendTableRow(
			a.FormatString(jobs[i].ID, &empty),
			a.FormatString(jobs[i].TrainingFile, &empty),
			a.FormatString(jobs[i].ValidationFile, &empty),
			a.FormatString(jobs[i].Model, &empty),
			a.FormatString(jobs[i].FineTunedModel, &empty),
			a.FormatTime(jobs[i].CreatedAt, &empty),
			a.FormatTime(jobs[i].FinishedAt, &empty),
			a.FormatString(jobs[i].Status, &empty),
			a.FormatExecStatus(jobs[i].ExecStatus),
		)
	}

	a.RenderTable()
	return nil
}
