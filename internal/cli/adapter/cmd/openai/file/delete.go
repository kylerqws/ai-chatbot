package file

import (
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctradp "github.com/kylerqws/chatbot/internal/cli/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

const allFlagKey = "all"

var filterFlagKeys = []string{
	allFlagKey,
	idFlagKey,
	statusFlagKey,
	purposeFlagKey,
	filenameFlagKey,
	createdAfterFlagKey,
	createdBeforeFlagKey,
}

type DeleteAdapter struct {
	*ListAdapter
}

func NewDeleteAdapter(app *intapp.App) ctradp.CommandAdapter {
	adp := &DeleteAdapter{}
	adp.ListAdapter = NewListAdapter(app).(*ListAdapter)

	return adp
}

func (a *DeleteAdapter) Configure() *cobra.Command {
	a.SetUse("delete <filter-flag> [filter-flag...]")
	a.SetShort("Delete one or more files from OpenAI account")
	a.SetLong("Use flags to filter files by multiple values.\n" +
		"You can repeat flags to provide more than one value, e.g.:\n  " + a.exampleString())

	a.SetFuncArgs(a.Validate)
	a.SetFuncRunE(a.Delete)

	a.ConfigureFlags()
	return a.MainConfigure()
}

func (a *DeleteAdapter) ConfigureFlags() {
	desc := "Delete all files (has highest priority over filters)\n"
	a.AddBoolFlag(allFlagKey, "", false, desc)

	a.ListAdapter.ConfigureFlags()
}

func (a *DeleteAdapter) Validate(cmd *cobra.Command, args []string) error {
	a.AddErrors(a.ValidateHasAnyFlags(filterFlagKeys...))
	if err := a.ListAdapter.Validate(cmd, args); err != nil {
		return err
	}

	return a.ErrorIfExist("One or more arguments/flags are invalid or missing.")
}

func (a *DeleteAdapter) Delete(_ *cobra.Command, _ []string) error {
	if a.Request() && a.ExistFiles() {
		if err := a.PrintFiles(); err != nil {
			return err
		}
	} else if !a.ExistErrors() {
		return a.PrintMessage("No files found.")
	}

	return a.ErrorIfExist("Failed to delete files or data is unavailable.")
}

func (a *DeleteAdapter) Request() bool {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().FileService()

	a.ListAdapter.Request()
	files := a.Files()

	for i := range files {
		resp, err := svc.DeleteFile(ctx, &ctrsvc.DeleteFileRequest{
			FileID: files[i].ID,
		})

		if err != nil {
			a.AddError(err)
		}
		files[i].ExecStatus = resp.Deleted
	}

	return true
}

func (a *DeleteAdapter) PrintFiles() error {
	_ = a.CreateTable()

	a.AppendTableHeader("File ID", "File Name", "Purpose",
		"Size", "Status", "Created", "State")

	a.SetColumnTableConfigs(
		a.ColumnConfig(1, text.AlignCenter, 27, text.Colors{text.Bold}),
		a.ColumnConfig(2, text.AlignRight, 19, nil),
		a.ColumnConfig(3, text.AlignRight, 19, nil),
		a.ColumnConfig(4, text.AlignRight, 10, nil),
		a.ColumnConfig(5, text.AlignRight, 10, nil),
		a.ColumnConfig(6, text.AlignRight, 19, nil),
		a.ColumnConfig(7, text.AlignCenter, 7, text.Colors{text.Bold}),
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
			a.FormatExecStatus(files[i].ExecStatus),
		)
	}

	a.RenderTable()
	return nil
}
