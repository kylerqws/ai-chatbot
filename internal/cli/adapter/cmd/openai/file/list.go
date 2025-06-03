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

type ListAdapter struct {
	*helper.CommandAdapter
	*helper.OpenAiAdapter
	*helper.OpenAiFileAdapter
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
	adp.OpenAiFileAdapter = helper.NewOpenAiFileAdapter(cmd)
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
	a.SetShort("Display files in OpenAI account")
	a.SetLong(a.HelpInfo())

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
		helper.PurposeFlagKey,
		helper.FilenameFlagKey,
		helper.CreatedAfterFlagKey,
		helper.CreatedBeforeFlagKey,
		helper.AfterFlagKey,
		helper.LimitFlagKey,
	)
}

func (a *ListAdapter) ConfigureFlags() {
	desc := "Filter by file ID (e.g. " + helper.FileIDExample + ")"
	a.AddStringSliceFlag(helper.IdFlagKey, "", []string{}, desc)

	desc = "Filter by status (" + a.FileStatusManager().JoinCodes(", ") + ")"
	a.AddStringSliceFlag(helper.StatusFlagKey, "", []string{}, desc)

	desc = "Filter by purpose (" + a.PurposeManager().JoinCodes(", ") + ")"
	a.AddStringSliceFlag(helper.PurposeFlagKey, "", []string{}, desc)

	desc = "Filter by file name (e.g. " + helper.Filename1Example + ", " + helper.Filename2Example + "...)"
	a.AddStringSliceFlag(helper.FilenameFlagKey, "", []string{}, desc)

	desc = "Filter by creation date after (e.g. " + helper.DateExample + " or " + helper.DatetimeExample + ")"
	a.AddStringFlag(helper.CreatedAfterFlagKey, "", "", desc)

	desc = "Filter by creation date before (e.g. " + helper.DateExample + " or " + helper.DatetimeExample + ")"
	a.AddStringFlag(helper.CreatedBeforeFlagKey, "", "", desc)

	desc = "After file ID (e.g. " + helper.FileIDExample + ")"
	a.AddStringFlag(helper.AfterFlagKey, "", "", desc)

	desc = "Limit files (" + a.JoinLimits(", ") + ")"
	a.AddUint8Flag(helper.LimitFlagKey, "", helper.DefaultLimit, desc)
}

func (a *ListAdapter) Validate(_ *cobra.Command, _ []string) error {
	a.AddErrors(a.ValidateStringSliceFlag(helper.IdFlagKey, a.ValidateFileID)...)

	a.AddErrors(a.ValidateStringSliceFlag(helper.StatusFlagKey, a.ValidateFileStatusCode)...)
	a.AddErrors(a.ValidateStringSliceFlag(helper.PurposeFlagKey, a.ValidatePurposeCode)...)

	a.AddErrors(
		a.ValidateStringFlag(helper.CreatedAfterFlagKey, a.ValidateDateFormat),
		a.ValidateStringFlag(helper.CreatedBeforeFlagKey, a.ValidateDateFormat),

		a.ValidateStringFlag(helper.AfterFlagKey, a.ValidateFileID),
		a.ValidateUint8Flag(helper.LimitFlagKey, a.ValidateLimit),
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

	fileIDs, err := fgs.GetStringSlice(helper.IdFlagKey)
	if err != nil {
		a.AddError(err)
	}

	statuses, err := fgs.GetStringSlice(helper.StatusFlagKey)
	if err != nil {
		a.AddError(err)
	}

	purposes, err := fgs.GetStringSlice(helper.PurposeFlagKey)
	if err != nil {
		a.AddError(err)
	}

	filenames, err := fgs.GetStringSlice(helper.FilenameFlagKey)
	if err != nil {
		a.AddError(err)
	}

	afterStr, err := fgs.GetString(helper.CreatedAfterFlagKey)
	if err != nil {
		a.AddError(err)
	}

	beforeStr, err := fgs.GetString(helper.CreatedBeforeFlagKey)
	if err != nil {
		a.AddError(err)
	}

	afterFileID, err := fgs.GetString(helper.AfterFlagKey)
	if err != nil {
		a.AddError(err)
	}

	limitFiles, err := fgs.GetUint8(helper.LimitFlagKey)
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
		AfterFileID:   afterFileID,
		LimitFiles:    limitFiles,
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
		"Size", "Created", "Status")

	a.SetColumnTableConfigs(
		a.ColumnConfig(1, text.AlignCenter, 27, text.Colors{text.Bold}),
		a.ColumnConfig(2, text.AlignRight, 19, nil),
		a.ColumnConfig(3, text.AlignRight, 19, nil),
		a.ColumnConfig(4, text.AlignRight, 10, nil),
		a.ColumnConfig(5, text.AlignRight, 19, nil),
		a.ColumnConfig(6, text.AlignRight, 10, text.Colors{text.Bold}),
	)

	files := a.Files()
	empty := helper.EmptyTableColumn

	for i := range files {
		a.AppendTableRow(
			a.FormatString(files[i].ID, &empty),
			a.FormatString(files[i].Filename, &empty),
			a.FormatString(files[i].Purpose, &empty),
			a.FormatBytes(files[i].Bytes, &empty),
			a.FormatTime(files[i].CreatedAt, &empty),
			a.FormatString(files[i].Status, &empty),
		)
	}

	a.RenderTable()
	return nil
}

func (a *ListAdapter) HelpInfo() string {
	cmdName := a.Command().Name()
	prpManager := a.PurposeManager()

	return "You can repeat flags to provide more than one value, e.g.:\n" +
		fmt.Sprintf(
			"  %s --%s %s --%s %s --%s %s",
			cmdName,
			helper.PurposeFlagKey, prpManager.Codes.FineTune,
			helper.PurposeFlagKey, prpManager.Codes.Evals,
			helper.CreatedAfterFlagKey, a.Date(0, 0, -7),
		)
}
