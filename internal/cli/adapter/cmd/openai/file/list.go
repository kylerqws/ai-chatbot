package file

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// ListAdapter provides the implementation for the OpenAI file listing CLI adapter.
type ListAdapter struct {
	*helper.CommandAdapter
	*helper.OpenAiAdapter
	*helper.OpenAiFileAdapter
	*helper.PaginationAdapter
	*helper.FilterAdapter
	*helper.SortAdapter
	*helper.FlagAdapter
	*helper.ValidateAdapter
	*helper.DateTimeAdapter
	*helper.TableAdapter
	*helper.FormatAdapter
}

// NewListAdapter creates a new ListAdapter adapter.
func NewListAdapter(app *intapp.App) ctr.CommandAdapter {
	adp := &ListAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapter = helper.NewCommandAdapter(app, cmd)
	adp.OpenAiAdapter = helper.NewOpenAiAdapter(cmd)
	adp.OpenAiFileAdapter = helper.NewOpenAiFileAdapter(cmd)
	adp.PaginationAdapter = helper.NewPaginationAdapter(cmd)
	adp.FilterAdapter = helper.NewFilterAdapter(cmd)
	adp.SortAdapter = helper.NewSortAdapter(cmd)
	adp.FlagAdapter = helper.NewFlagAdapter(cmd)
	adp.ValidateAdapter = helper.NewValidateAdapter(cmd)
	adp.DateTimeAdapter = helper.NewDateTimeAdapter(cmd)
	adp.TableAdapter = helper.NewTableAdapter(cmd)
	adp.FormatAdapter = helper.NewFormatAdapter(cmd)

	return adp
}

// Configure applies configuration for the command.
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

// HelpInfo returns extended help usage text for the command.
func (a *ListAdapter) HelpInfo() string {
	return "You can repeat flags to provide more than one filter, for example:\n" +
		fmt.Sprintf(
			"  %s --%s '%s' --%s '%s' --%s '%s'",
			a.Command().Name(),
			helper.PurposeFlagKey, a.PurposeManager().Codes.FineTune,
			helper.PurposeFlagKey, helper.PurposeExample,
			helper.SortOrderFlagKey, helper.SortOrderDescExample,
		)
}

// ConfigureFilters defines available filter keys for the command.
func (a *ListAdapter) ConfigureFilters() {
	a.AddFilterKeys(
		helper.IdFlagKey,
		helper.PurposeFlagKey,
		helper.FilenameFlagKey,
		helper.CreatedAfterFlagKey,
		helper.CreatedBeforeFlagKey,
		helper.ExpiresAfterFlagKey,
		helper.ExpiresBeforeFlagKey,
		helper.SortOrderFlagKey,
		helper.AfterFlagKey,
		helper.LimitFlagKey,
	)
}

// ConfigureFlags registers available flags for the command.
func (a *ListAdapter) ConfigureFlags() {
	desc := "Filter by file ID (e.g. " + helper.FileIDExample + ")"
	a.AddStringSliceFlag(helper.IdFlagKey, "", []string{}, desc)

	desc = "Filter by purpose (e.g. " + a.PurposeManager().JoinCodes(", ") + ", ...)"
	a.AddStringSliceFlag(helper.PurposeFlagKey, "", []string{}, desc)

	desc = "Filter by file name (e.g. " + helper.Filename1Example + ", " + helper.Filename2Example + ", ...)\n"
	a.AddStringSliceFlag(helper.FilenameFlagKey, "", []string{}, desc)

	desc = "Filter by creation date after (e.g. " + helper.DateExample + " or '" + helper.DatetimeExample + "')"
	a.AddStringFlag(helper.CreatedAfterFlagKey, "", "", desc)

	desc = "Filter by creation date before (e.g. " + helper.DateExample + " or '" + helper.DatetimeExample + "')"
	a.AddStringFlag(helper.CreatedBeforeFlagKey, "", "", desc)

	desc = "Filter by expires date after (e.g. " + helper.DateExample + " or '" + helper.DatetimeExample + "')"
	a.AddStringFlag(helper.ExpiresAfterFlagKey, "", "", desc)

	desc = "Filter by expires date before (e.g. " + helper.DateExample + " or '" + helper.DatetimeExample + "')\n"
	a.AddStringFlag(helper.ExpiresBeforeFlagKey, "", "", desc)

	desc = "Sort Order (val. " + helper.SortOrderAscExample + " or " + helper.SortOrderDescExample + ")"
	a.AddStringFlag(helper.SortOrderFlagKey, "", helper.DefaultSortOrder, desc)

	desc = "After file ID (e.g. " + helper.FileIDExample + ")"
	a.AddStringFlag(helper.AfterFlagKey, "", "", desc)

	desc = "Limit files (" + a.JoinLimits(", ") + ")"
	a.AddUint8Flag(helper.LimitFlagKey, "", helper.DefaultLimit, desc)
}

// Validate validates all arguments and flags passed to the command.
func (a *ListAdapter) Validate(_ *cobra.Command, _ []string) error {
	a.AddErrors(a.ValidateStringSliceFlag(helper.IdFlagKey, a.ValidateFileID)...)
	a.AddErrors(a.ValidateStringSliceFlag(helper.PurposeFlagKey, a.ValidatePurposeCode)...)

	a.AddErrors(
		a.ValidateStringFlag(helper.CreatedAfterFlagKey, a.ValidateDateFormat),
		a.ValidateStringFlag(helper.CreatedBeforeFlagKey, a.ValidateDateFormat),

		a.ValidateStringFlag(helper.ExpiresAfterFlagKey, a.ValidateDateFormat),
		a.ValidateStringFlag(helper.ExpiresBeforeFlagKey, a.ValidateDateFormat),

		a.ValidateStringFlag(helper.SortOrderFlagKey, a.ValidateSortOrder),
		a.ValidateStringFlag(helper.AfterFlagKey, a.ValidateFileID),
		a.ValidateUint8Flag(helper.LimitFlagKey, a.ValidateLimit),
	)

	return a.ErrorIfExist("One or more arguments/flags are invalid or missing.")
}

// List executes the file listing process.
func (a *ListAdapter) List(_ *cobra.Command, _ []string) error {
	a.Request()

	if a.ExistFiles() {
		if err := a.PrintFiles(); err != nil {
			return err
		}
	} else if !a.ExistErrors() {
		return a.PrintMessage("No files found.")
	}

	return a.ErrorIfExist("Failed to retrieve files or data is unavailable.")
}

// Request executes the API call to retrieve files from OpenAI.
func (a *ListAdapter) Request() {
	app := a.App()
	ctx := app.Context()

	svc := app.OpenAI().ServiceProvider().File()
	req := svc.NewListFilesRequest()

	fileIDs, err := a.GetStringSliceFlag(helper.IdFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.FileIDs = fileIDs

	purposes, err := a.GetStringSliceFlag(helper.PurposeFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.Purpose, req.Purposes = a.ExtractSinglePurpose(&purposes)

	filenames, err := a.GetStringSliceFlag(helper.FilenameFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.Filenames = filenames

	createdAfter, err := a.GetStringFlag(helper.CreatedAfterFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.CreatedAfter = a.ParseDateTime(createdAfter)

	createdBefore, err := a.GetStringFlag(helper.CreatedBeforeFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.CreatedBefore = a.ParseDateTime(createdBefore)

	expiresAfter, err := a.GetStringFlag(helper.ExpiresAfterFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.ExpiresAfter = a.ParseDateTime(expiresAfter)

	expiresBefore, err := a.GetStringFlag(helper.ExpiresBeforeFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.ExpiresBefore = a.ParseDateTime(expiresBefore)

	order, err := a.GetPointerStringFlag(helper.SortOrderFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.Order = a.SortOrderToLower(order)

	afterID, err := a.GetPointerStringFlag(helper.AfterFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.After = afterID

	limit, err := a.GetPointerUint8Flag(helper.LimitFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.Limit = limit

	resp, err := svc.ListFiles(ctx, req)
	if err != nil {
		a.AddError(err)
	}
	a.AddFiles(a.WrapOpenAIFiles(resp.Files...)...)
}

// PrintFiles renders the retrieved files in a formatted table.
func (a *ListAdapter) PrintFiles() error {
	_ = a.CreateTable()
	a.AppendTableHeader("File ID", "Purpose", "File Name", "Size", "Created", "Expires")

	a.SetColumnTableConfigs(
		a.ColumnConfig(1, text.AlignLeft, 27, text.Colors{text.Bold}),
		a.ColumnConfig(2, text.AlignRight, 19, nil),
		a.ColumnConfig(3, text.AlignRight, 19, nil),
		a.ColumnConfig(4, text.AlignRight, 10, nil),
		a.ColumnConfig(5, text.AlignRight, 19, nil),
		a.ColumnConfig(6, text.AlignRight, 19, nil),
	)

	files := a.Files()
	empty := helper.EmptyTableColumn

	for i := range files {
		a.AppendTableRow(
			a.FormatString(&files[i].ID, &empty),
			a.FormatString(&files[i].Purpose, &empty),
			a.FormatString(&files[i].Filename, &empty),
			a.FormatBytes(&files[i].Bytes, &empty),
			a.FormatTime(&files[i].CreatedAt, &empty),
			a.FormatTime(files[i].ExpiresAt, &empty),
		)
	}

	a.RenderTable()
	return nil
}

// ExtractSinglePurpose returns the single purpose if only one exists.
func (*ListAdapter) ExtractSinglePurpose(purposes *[]string) (*string, []string) {
	if purposes == nil || len(*purposes) == 0 {
		return nil, nil
	}
	if len(*purposes) == 1 {
		return &(*purposes)[0], nil
	}
	return nil, *purposes
}
