package file

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

// ListAdapter provides the implementation for the list CLI adapter.
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

// NewListAdapter creates a new list command adapter.
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
	return "You can repeat flags to provide more than one filter, e.g.:\n" +
		fmt.Sprintf(
			"  %s --%s %s --%s %s --%s %s",
			a.Command().Name(),
			helper.PurposeFlagKey, a.PurposeManager().Codes.FineTune,
			helper.PurposeFlagKey, helper.PurposeExample,
			helper.SortOrderFlagKey, helper.SortDesc,
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

	desc = "Filter by purpose (e.g. " + a.PurposeManager().JoinCodes(", ") + "...)"
	a.AddStringSliceFlag(helper.PurposeFlagKey, "", []string{}, desc)

	desc = "Filter by file name (e.g. " + helper.Filename1Example + ", " + helper.Filename2Example + "...)"
	a.AddStringSliceFlag(helper.FilenameFlagKey, "", []string{}, desc)

	desc = "Filter by creation date after (e.g. " + helper.DateExample + " or " + helper.DatetimeExample + ")"
	a.AddStringFlag(helper.CreatedAfterFlagKey, "", "", desc)

	desc = "Filter by creation date before (e.g. " + helper.DateExample + " or " + helper.DatetimeExample + ")"
	a.AddStringFlag(helper.CreatedBeforeFlagKey, "", "", desc)

	desc = "Filter by expires date after (e.g. " + helper.DateExample + " or " + helper.DatetimeExample + ")"
	a.AddStringFlag(helper.ExpiresAfterFlagKey, "", "", desc)

	desc = "Filter by expires date before (e.g. " + helper.DateExample + " or " + helper.DatetimeExample + ")"
	a.AddStringFlag(helper.ExpiresBeforeFlagKey, "", "", desc)

	desc = "Sort order for list (" + helper.SortAsc + " or " + helper.SortDesc + ")"
	a.AddStringFlag(helper.SortOrderFlagKey, "", helper.DefaultSort, desc)

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
	if a.Request() && a.ExistFiles() {
		if err := a.PrintFiles(); err != nil {
			return err
		}
	} else if !a.ExistErrors() {
		return a.PrintMessage("No files found.")
	}

	return a.ErrorIfExist("Failed to retrieve files or data is unavailable.")
}

// Request executes the API call to retrieve files from OpenAI.
func (a *ListAdapter) Request() bool {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().ServiceProvider().File()
	fgs := a.Command().Flags()
	req := svc.NewListFilesRequest()

	// File IDs
	fileIDs, err := fgs.GetStringSlice(helper.IdFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.FileIDs = fileIDs

	// Purposes
	purposes, err := fgs.GetStringSlice(helper.PurposeFlagKey)
	if err != nil {
		a.AddError(err)
	}
	purpose, purposeList := a.extractSingleValue(purposes)
	req.Purposes = purposeList
	req.Purpose = &purpose

	// Filenames
	filenames, err := fgs.GetStringSlice(helper.FilenameFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.Filenames = filenames

	// Created After
	createdAfter, err := fgs.GetString(helper.CreatedAfterFlagKey)
	if err != nil {
		a.AddError(err)
	}
	if createdAfter != "" {
		req.CreatedAfter = a.ParseDateTime(createdAfter)
	}

	// Created Before
	createdBefore, err := fgs.GetString(helper.CreatedBeforeFlagKey)
	if err != nil {
		a.AddError(err)
	}
	if createdBefore != "" {
		req.CreatedBefore = a.ParseDateTime(createdBefore)
	}

	// Expires After
	expiresAfter, err := fgs.GetString(helper.ExpiresAfterFlagKey)
	if err != nil {
		a.AddError(err)
	}
	if expiresAfter != "" {
		req.ExpiresAfter = a.ParseDateTime(expiresAfter)
	}

	// Expires Before
	expiresBefore, err := fgs.GetString(helper.ExpiresBeforeFlagKey)
	if err != nil {
		a.AddError(err)
	}
	if expiresBefore != "" {
		req.ExpiresBefore = a.ParseDateTime(expiresBefore)
	}

	// Sort Order
	sortOrder, err := fgs.GetString(helper.SortOrderFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.Order = &sortOrder

	// After File ID
	afterID, err := fgs.GetString(helper.AfterFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.After = &afterID

	// Limit
	limit, err := fgs.GetUint8(helper.LimitFlagKey)
	if err != nil {
		a.AddError(err)
	}
	req.Limit = &limit

	// API call
	resp, err := svc.ListFiles(ctx, req)
	if err != nil {
		a.AddError(err)
	}

	a.AddFiles(a.WrapOpenAIFiles(resp.Files...)...)
	return true
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

// extractSingleValue returns the single value if only one exists.
func (*ListAdapter) extractSingleValue(values []string) (string, []string) {
	if len(values) == 1 {
		return values[0], nil
	}
	return "", values
}
