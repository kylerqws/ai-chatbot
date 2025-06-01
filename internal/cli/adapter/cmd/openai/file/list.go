package file

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctradp "github.com/kylerqws/chatbot/internal/cli/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

const (
	idFlagKey            = "id"
	statusFlagKey        = "status"
	purposeFlagKey       = "purpose"
	filenameFlagKey      = "filename"
	createdAfterFlagKey  = "created-after"
	createdBeforeFlagKey = "created-before"
)

const (
	fileIDExample    = "file-xxxxxx..."
	filename1Example = "step_metrics.csv"
	filename2Example = "prompts.json"
	dateExample      = "1970-01-01"
	datetimeExample  = "1970-01-01 00:00:00"
)

type ListAdapter struct {
	*helper.CommandAdapter
	*helper.OpenAiAdapter
	*helper.OpenAiFileAdapter
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
	adp.OpenAiFileAdapter = helper.NewOpenAiFileAdapter(cmd)
	adp.FlagAdapter = helper.NewFlagAdapter(cmd)
	adp.ValidateAdapter = helper.NewValidateAdapter(cmd)
	adp.DateTimeAdapter = helper.NewDateTimeAdapter(cmd)
	adp.TableAdapter = helper.NewTableAdapter(cmd)
	adp.FormatAdapter = helper.NewFormatAdapter(cmd)

	return adp
}

func (a *ListAdapter) Configure() *cobra.Command {
	a.SetUse("list")
	a.SetShort("Display files in OpenAI account")
	a.SetLong("Repeat flags to filter by multiple values, e.g.:\n  " + a.exampleString())

	a.SetFuncArgs(a.Validate)
	a.SetFuncRunE(a.List)

	a.ConfigureFlags()
	return a.MainConfigure()
}

func (a *ListAdapter) ConfigureFlags() {
	desc := "Filter by file ID (e.g. " + fileIDExample + ")"
	a.AddStringSliceFlag(idFlagKey, "", []string{}, desc)

	desc = "Filter by status (e.g. " + a.FileStatusManager().JoinCodes(", ") + ")"
	a.AddStringSliceFlag(statusFlagKey, "", []string{}, desc)

	desc = "Filter by purpose (e.g. " + a.PurposeManager().JoinCodes(", ") + ")"
	a.AddStringSliceFlag(purposeFlagKey, "", []string{}, desc)

	desc = "Filter by file name (e.g. " + filename1Example + ", " + filename2Example + ")"
	a.AddStringSliceFlag(filenameFlagKey, "", []string{}, desc)

	desc = "Filter by creation date after (e.g. " + dateExample + " or " + datetimeExample + ")"
	a.AddStringFlag(createdAfterFlagKey, "", a.DateTime(0, -3, 0), desc)

	desc = "Filter by creation date before (e.g. " + dateExample + " or " + datetimeExample + ")"
	a.AddStringFlag(createdBeforeFlagKey, "", "", desc)
}

func (a *ListAdapter) Validate(_ *cobra.Command, _ []string) error {
	a.AddErrors(a.ValidateStringSliceFlag(idFlagKey, a.ValidateFileID)...)

	a.AddErrors(a.ValidateStringSliceFlag(statusFlagKey, a.ValidateFileStatusCode)...)
	a.AddErrors(a.ValidateStringSliceFlag(purposeFlagKey, a.ValidatePurposeCode)...)

	a.AddErrors(
		a.ValidateStringFlag(createdAfterFlagKey, a.ValidateDateFormat),
		a.ValidateStringFlag(createdBeforeFlagKey, a.ValidateDateFormat),
	)

	return a.ErrorIfExist("One or more arguments/flags are invalid or missing.")
}

func (a *ListAdapter) List(_ *cobra.Command, _ []string) error {
	if a.Request() && a.ExistFiles() {
		if err := a.PrintFiles(); err != nil {
			return err
		}
	} else if !a.ExistErrors() {
		return a.PrintMessage("No files found.")
	}

	return a.ErrorIfExist("Failed to retrieve files or data is unavailable.")
}

func (a *ListAdapter) Request() bool {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().FileService()
	fgs := a.Command().Flags()

	fileIDs, err := fgs.GetStringSlice(idFlagKey)
	if err != nil {
		a.AddError(err)
	}

	statuses, err := fgs.GetStringSlice(statusFlagKey)
	if err != nil {
		a.AddError(err)
	}

	purposes, err := fgs.GetStringSlice(purposeFlagKey)
	if err != nil {
		a.AddError(err)
	}

	filenames, err := fgs.GetStringSlice(filenameFlagKey)
	if err != nil {
		a.AddError(err)
	}

	afterStr, err := fgs.GetString(createdAfterFlagKey)
	if err != nil {
		a.AddError(err)
	}

	beforeStr, err := fgs.GetString(createdBeforeFlagKey)
	if err != nil {
		a.AddError(err)
	}

	resp, err := svc.ListFiles(ctx, &ctrsvc.ListFilesRequest{
		FileIDs:       fileIDs,
		Statuses:      statuses,
		Purposes:      purposes,
		Filenames:     filenames,
		CreatedAfter:  a.ParseDateTime(afterStr),
		CreatedBefore: a.ParseDateTime(beforeStr),
	})

	if err != nil {
		a.AddError(err)
	}
	a.AddFiles(a.WrapOpenAIFiles(resp.Files...)...)

	return true
}

func (a *ListAdapter) PrintFiles() error {
	_ = a.CreateTable()

	a.AppendTableHeader("File ID", "File Name", "Purpose",
		"Size", "Status", "Created")

	a.SetColumnTableConfigs(
		a.ColumnConfig(1, text.AlignCenter, 27, text.Colors{text.Bold}),
		a.ColumnConfig(2, text.AlignRight, 19, nil),
		a.ColumnConfig(3, text.AlignRight, 19, nil),
		a.ColumnConfig(4, text.AlignRight, 10, nil),
		a.ColumnConfig(5, text.AlignRight, 10, nil),
		a.ColumnConfig(6, text.AlignRight, 19, nil),
	)

	files := a.Files()
	empty := helper.EmptyTableColumn

	for i := range files {
		a.AppendTableRow(
			a.FormatString(files[i].ID, &empty),
			a.FormatString(files[i].Filename, &empty),
			a.FormatString(files[i].Purpose, &empty),
			a.FormatBytes(files[i].Bytes, &empty),
			a.FormatString(files[i].Status, &empty),
			a.FormatTime(files[i].CreatedAt, &empty),
		)
	}

	a.RenderTable()
	return nil
}

func (a *ListAdapter) exampleString() string {
	cmdName := a.Command().Name()
	prpManager := a.PurposeManager()

	return fmt.Sprintf("%s --%s %s --%s %s --%s %s", cmdName,
		purposeFlagKey, prpManager.Codes.FineTune, purposeFlagKey, prpManager.Codes.Evals,
		createdAfterFlagKey, a.Date(0, -3, 0))
}
