package job

import (
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctradp "github.com/kylerqws/chatbot/internal/cli/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type ListAdapter struct {
	*helper.CommandAdapter
	*helper.OpenAiAdapter
	*helper.OpenAiJobAdapter
	*helper.PaginationAdapter
	*helper.FilterAdapter
	*helper.FlagAdapter
	*helper.ValidateAdapter
	*helper.DateTimeAdapter
	*helper.TableAdapter
	*helper.FormatAdapter
}

func NewListAdapter(app *intapp.App) ctradp.CommandAdapter {
	adp := &ListAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapter = helper.NewCommandAdapter(app, cmd)
	adp.OpenAiAdapter = helper.NewOpenAiAdapter(cmd)
	adp.OpenAiJobAdapter = helper.NewOpenAiJobAdapter(cmd)
	adp.PaginationAdapter = helper.NewPaginationAdapter(cmd)
	adp.FilterAdapter = helper.NewFilterAdapter(cmd)
	adp.FlagAdapter = helper.NewFlagAdapter(cmd)
	adp.ValidateAdapter = helper.NewValidateAdapter(cmd)
	adp.DateTimeAdapter = helper.NewDateTimeAdapter(cmd)
	adp.TableAdapter = helper.NewTableAdapter(cmd)
	adp.FormatAdapter = helper.NewFormatAdapter(cmd)

	return adp
}

func (a *ListAdapter) Configure() *cobra.Command {
	a.SetUse("list [filter-flag...]")
	a.SetShort("Display jobs in OpenAI account")
	a.SetLong(a.JobListHelpInfo(a.OpenAiAdapter))

	a.SetFuncArgs(a.Validate)
	a.SetFuncRunE(a.List)

	a.ConfigureFilters()
	a.ConfigureFlags()

	return a.MainConfigure()
}

func (a *ListAdapter) ConfigureFilters() {
	a.AddFilterKeys(
		helper.IdFlagKey,
		helper.StatusFlagKey,
		helper.ModelFlagKey,
		helper.FineTunedModelFlagKey,
		helper.TrainingFileFlagKey,
		helper.ValidationFileFlagKey,
		helper.CreatedAfterFlagKey,
		helper.CreatedBeforeFlagKey,
		helper.FinishedAfterFlagKey,
		helper.FinishedBeforeFlagKey,
		helper.AfterFlagKey,
		helper.LimitFlagKey,
	)
}

func (a *ListAdapter) ConfigureFlags() {
	desc := "Filter by job ID (e.g. " + helper.JobIDExample + ")"
	a.AddStringSliceFlag(helper.IdFlagKey, "", []string{}, desc)

	desc = "Filter by status (" + a.JobStatusManager().JoinCodes(", ") + ")"
	a.AddStringSliceFlag(helper.StatusFlagKey, "", []string{}, desc)

	desc = "Filter by model (e.g. " + a.ModelManager().JoinCodes(", ") + "...)"
	a.AddStringSliceFlag(helper.ModelFlagKey, "", []string{}, desc)

	desc = "Filter by fine-tuned model (e.g. " + helper.FineTunedModelExample + ")"
	a.AddStringSliceFlag(helper.FineTunedModelFlagKey, "", []string{}, desc)

	desc = "Filter by training file ID (e.g. " + helper.FileIDExample + ")"
	a.AddStringSliceFlag(helper.TrainingFileFlagKey, "", []string{}, desc)

	desc = "Filter by validation file ID (e.g. " + helper.FileIDExample + ")"
	a.AddStringSliceFlag(helper.ValidationFileFlagKey, "", []string{}, desc)

	desc = "Filter by creation date after (e.g. " + helper.DateExample + " or " + helper.DatetimeExample + ")"
	a.AddStringFlag(helper.CreatedAfterFlagKey, "", "", desc)

	desc = "Filter by creation date before (e.g. " + helper.DateExample + " or " + helper.DatetimeExample + ")"
	a.AddStringFlag(helper.CreatedBeforeFlagKey, "", "", desc)

	desc = "Filter by completion date after (e.g. " + helper.DateExample + " or " + helper.DatetimeExample + ")"
	a.AddStringFlag(helper.FinishedAfterFlagKey, "", "", desc)

	desc = "Filter by completion date before (e.g. " + helper.DateExample + " or " + helper.DatetimeExample + ")"
	a.AddStringFlag(helper.FinishedBeforeFlagKey, "", "", desc)

	desc = "After job ID (e.g. " + helper.JobIDExample + ")"
	a.AddStringFlag(helper.AfterFlagKey, "", "", desc)

	desc = "Limit jobs (" + a.JoinLimits(", ") + ")"
	a.AddUint8Flag(helper.LimitFlagKey, "", helper.DefaultLimit, desc)
}

func (a *ListAdapter) Validate(_ *cobra.Command, _ []string) error {
	a.AddErrors(a.ValidateStringSliceFlag(helper.IdFlagKey, a.ValidateJobID)...)
	a.AddErrors(a.ValidateStringSliceFlag(helper.StatusFlagKey, a.ValidateJobStatusCode)...)

	a.AddErrors(a.ValidateStringSliceFlag(helper.ModelFlagKey, a.ValidateModelCode)...)
	a.AddErrors(a.ValidateStringSliceFlag(helper.FineTunedModelFlagKey, a.ValidateModelCode)...)

	a.AddErrors(a.ValidateStringSliceFlag(helper.TrainingFileFlagKey, a.ValidateFileID)...)
	a.AddErrors(a.ValidateStringSliceFlag(helper.ValidationFileFlagKey, a.ValidateFileID)...)

	a.AddErrors(
		a.ValidateStringFlag(helper.CreatedAfterFlagKey, a.ValidateDateFormat),
		a.ValidateStringFlag(helper.CreatedBeforeFlagKey, a.ValidateDateFormat),

		a.ValidateStringFlag(helper.FinishedAfterFlagKey, a.ValidateDateFormat),
		a.ValidateStringFlag(helper.FinishedBeforeFlagKey, a.ValidateDateFormat),

		a.ValidateStringFlag(helper.AfterFlagKey, a.ValidateJobID),
		a.ValidateUint8Flag(helper.LimitFlagKey, a.ValidateLimit),
	)

	return a.ErrorIfExist("One or more arguments/flags are invalid or missing.")
}

func (a *ListAdapter) List(_ *cobra.Command, _ []string) error {
	if a.Request() && a.ExistJobs() {
		if err := a.PrintJobs(); err != nil {
			return err
		}
	} else if !a.ExistErrors() {
		return a.PrintMessage("No jobs found.")
	}

	return a.ErrorIfExist("Failed to retrieve jobs or data is unavailable.")
}

func (a *ListAdapter) Request() bool {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().JobService()
	fgs := a.Command().Flags()

	jobIDs, err := fgs.GetStringSlice(helper.IdFlagKey)
	if err != nil {
		a.AddError(err)
	}

	statuses, err := fgs.GetStringSlice(helper.StatusFlagKey)
	if err != nil {
		a.AddError(err)
	}

	models, err := fgs.GetStringSlice(helper.ModelFlagKey)
	if err != nil {
		a.AddError(err)
	}

	fineTunedModels, err := fgs.GetStringSlice(helper.FineTunedModelFlagKey)
	if err != nil {
		a.AddError(err)
	}

	trainingFileIDs, err := fgs.GetStringSlice(helper.TrainingFileFlagKey)
	if err != nil {
		a.AddError(err)
	}

	validationFileIDs, err := fgs.GetStringSlice(helper.ValidationFileFlagKey)
	if err != nil {
		a.AddError(err)
	}

	createdAfterStr, err := fgs.GetString(helper.CreatedAfterFlagKey)
	if err != nil {
		a.AddError(err)
	}

	createdBeforeStr, err := fgs.GetString(helper.CreatedBeforeFlagKey)
	if err != nil {
		a.AddError(err)
	}

	finishedAfterStr, err := fgs.GetString(helper.FinishedAfterFlagKey)
	if err != nil {
		a.AddError(err)
	}

	finishedBeforeStr, err := fgs.GetString(helper.FinishedBeforeFlagKey)
	if err != nil {
		a.AddError(err)
	}

	afterJobID, err := fgs.GetString(helper.AfterFlagKey)
	if err != nil {
		a.AddError(err)
	}

	limitJobs, err := fgs.GetUint8(helper.LimitFlagKey)
	if err != nil {
		a.AddError(err)
	}

	resp, err := svc.ListJobs(ctx, &ctrsvc.ListJobsRequest{
		JobIDs:          jobIDs,
		Statuses:        statuses,
		Models:          models,
		FineTunedModels: fineTunedModels,
		TrainingFiles:   trainingFileIDs,
		ValidationFiles: validationFileIDs,
		CreatedAfter:    a.ParseDateTime(createdAfterStr),
		CreatedBefore:   a.ParseDateTime(createdBeforeStr),
		FinishedAfter:   a.ParseDateTime(finishedAfterStr),
		FinishedBefore:  a.ParseDateTime(finishedBeforeStr),
		AfterJobID:      afterJobID,
		LimitJobs:       limitJobs,
	})

	if err != nil {
		a.AddError(err)
	}
	a.AddJobs(a.WrapOpenAIJobs(resp.Jobs...)...)

	return true
}

func (a *ListAdapter) PrintJobs() error {
	_ = a.CreateTable()

	a.AppendTableHeader("Job ID", "Training File", "Validation File",
		"Model", "Fine-Tuned Model", "Created", "Finished", "Status")

	a.SetColumnTableConfigs(
		a.ColumnConfig(1, text.AlignCenter, 27, text.Colors{text.Bold}),
		a.ColumnConfig(2, text.AlignRight, 19, nil),
		a.ColumnConfig(3, text.AlignRight, 19, nil),
		a.ColumnConfig(4, text.AlignRight, 19, nil),
		a.ColumnConfig(5, text.AlignRight, 19, nil),
		a.ColumnConfig(6, text.AlignRight, 19, nil),
		a.ColumnConfig(7, text.AlignRight, 19, nil),
		a.ColumnConfig(8, text.AlignRight, 10, nil),
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
		)
	}

	a.RenderTable()
	return nil
}
