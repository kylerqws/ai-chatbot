package job

import (
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctradp "github.com/kylerqws/chatbot/internal/cli/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type CreateAdapter struct {
	*helper.CommandAdapter
	*helper.OpenAiAdapter
	*helper.OpenAiJobAdapter
	*helper.OpenAiFileAdapter
	*helper.ArgumentAdapter
	*helper.FlagAdapter
	*helper.ValidateAdapter
	*helper.TableAdapter
	*helper.FormatAdapter
}

func NewCreateAdapter(app *intapp.App) ctradp.CommandAdapter {
	adp := &CreateAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapter = helper.NewCommandAdapter(app, cmd)
	adp.OpenAiAdapter = helper.NewOpenAiAdapter(cmd)
	adp.OpenAiJobAdapter = helper.NewOpenAiJobAdapter(cmd)
	adp.OpenAiFileAdapter = helper.NewOpenAiFileAdapter(cmd)
	adp.ArgumentAdapter = helper.NewArgumentAdapter(cmd)
	adp.FlagAdapter = helper.NewFlagAdapter(cmd)
	adp.ValidateAdapter = helper.NewValidateAdapter(cmd)
	adp.TableAdapter = helper.NewTableAdapter(cmd)
	adp.FormatAdapter = helper.NewFormatAdapter(cmd)

	return adp
}

func (a *CreateAdapter) Configure() *cobra.Command {
	argExample := "training-file-id[:validation-file-id][:model]"
	a.SetUse("create <" + argExample + "> [" + argExample + "...]")
	a.SetShort("Create one or more jobs in OpenAI account")
	a.SetLong(a.JobCreateHelpInfo(a.OpenAiAdapter))

	a.SetFuncArgs(a.Validate)
	a.SetFuncRunE(a.Create)

	a.ConfigureFlags()
	return a.MainConfigure()
}

func (a *CreateAdapter) ConfigureFlags() {
	mdlManager := a.ModelManager()

	desc := "Default model for training files without ':model' suffix " +
		"(" + mdlManager.JoinCodes(", ") + ")"
	a.AddStringFlag(helper.DefaultModelFlagKey, "", mdlManager.Default().Code, desc)

	desc = "Default validation file ID for training files without ':validation-file-id' suffix " +
		"(e.g. " + helper.FileIDExample + ")"
	a.AddStringFlag(helper.DefaultValidationFileFlagKey, "", "", desc)
}

func (a *CreateAdapter) Validate(_ *cobra.Command, _ []string) error {
	a.AddError(a.ValidateHasMoreArgsThan(0))

	a.AddErrors(a.ValidateArgs()...)
	a.AddErrors(a.ValidateFlags()...)

	return a.ErrorIfExist("One or more arguments/flags are invalid or missing.")
}

func (a *CreateAdapter) ValidateArgs() []error {
	var errs []error
	args := a.Command().Flags().Args()

	for i := range args {
		tFileID, vFileID, mdlCode := a.SplitArg(args[i])

		err := a.ValidateFileID(tFileID)
		if err != nil {
			errs = append(errs, err)
		}

		if mdlCode != "" {
			err = a.ValidateModelCode(mdlCode)
			if err != nil {
				errs = append(errs, err)
			}
		}

		if vFileID != "" {
			err = a.ValidateFileID(vFileID)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	return errs
}

func (a *CreateAdapter) ValidateFlags() []error {
	var errs []error

	err := a.ValidateStringFlag(helper.DefaultModelFlagKey, a.ValidateModelCode)
	if err != nil {
		errs = append(errs, err)
	}

	err = a.ValidateStringFlag(helper.DefaultValidationFileFlagKey, a.ValidateFileID)
	if err != nil {
		errs = append(errs, err)
	}

	return errs
}

func (a *CreateAdapter) Create(_ *cobra.Command, _ []string) error {
	if a.Request() && a.ExistJobs() {
		if err := a.PrintJobs(); err != nil {
			return err
		}
	}

	return a.ErrorIfExist("Failed to create jobs or data is unavailable.")
}

func (a *CreateAdapter) Request() bool {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().JobService()
	fgs := a.Command().Flags()
	args := fgs.Args()

	defaultMdlCode, err := fgs.GetString(helper.DefaultModelFlagKey)
	if err != nil {
		a.AddError(err)
	}

	defaultVFileID, err := fgs.GetString(helper.DefaultValidationFileFlagKey)
	if err != nil {
		a.AddError(err)
	}

	for i := range args {
		tFileID, vFileID, mdlCode :=
			a.SplitArg(args[i], defaultVFileID, defaultMdlCode)

		resp, err := svc.CreateJob(ctx, &ctrsvc.CreateJobRequest{
			Model:          mdlCode,
			TrainingFile:   tFileID,
			ValidationFile: vFileID,
		})

		if err != nil {
			a.AddError(err)

			resp.Job = &ctrsvc.Job{
				Model:          mdlCode,
				TrainingFile:   tFileID,
				ValidationFile: vFileID,
			}
		}

		job := a.WrapOpenAIJob(resp.Job)
		job.ExecStatus = err == nil

		a.AddJob(job)
	}

	return true
}

func (a *CreateAdapter) PrintJobs() error {
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
		a.ColumnConfig(8, text.AlignRight, 10, text.Colors{text.Bold}),
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

func (a *CreateAdapter) SplitArg(arg string, defaults ...string) (string, string, string) {
	tFileID, vFileID, mdlCode := a.SplitTripleArg(arg)
	defCount := len(defaults)

	if vFileID != "" && mdlCode == "" {
		_, err := a.ModelManager().Resolve(vFileID)
		if err == nil {
			mdlCode, vFileID = vFileID, ""
		}
	}

	if vFileID == "" && defCount > 0 {
		vFileID = defaults[0]
	}
	if mdlCode == "" && defCount > 1 {
		mdlCode = defaults[1]
	}

	return tFileID, vFileID, mdlCode
}
