package file

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	intapp "github.com/kylerqws/chatbot/internal/app"
	inthlp "github.com/kylerqws/chatbot/internal/cli/helper"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

const (
	allFlagKey = "all"
)

var (
	allFlagKeys = []string{
		allFlagKey,
		idFlagKey,
		purposeFlagKey,
		createdAfterFlagKey,
		createdBeforeFlagKey,
	}
)

type DeleteAdapter struct {
	*ListAdapter
}

func NewDeleteAdapter(app *intapp.App) ctr.CommandAdapter {
	adp := &DeleteAdapter{}
	adp.ListAdapter = NewListAdapter(app).(*ListAdapter)

	return adp
}

func (a *DeleteAdapter) Configure() *cobra.Command {
	a.SetUse("delete <flag>")
	a.SetShort("Delete one or more files by filter from OpenAI")

	a.SetFuncArgs(a.FuncArgs)
	a.SetFuncRunE(a.FuncRunE)

	a.AddFlags()
	return a.MainConfigure()
}

func (a *DeleteAdapter) AddFlags() {
	var desc string

	desc = "Delete all files in your OpenAI account\n"
	a.AddBoolFlag(allFlagKey, "", false, desc)

	a.ListAdapter.AddFlags()
}

func (a *DeleteAdapter) FuncArgs(cmd *cobra.Command, _ []string) error {
	if !a.HasAnyFlag(allFlagKeys...) {
		return fmt.Errorf("at least one filter flag must be specified, usage: %s", cmd.UseLine())
	}
	return nil
}

func (a *DeleteAdapter) FuncRunE(_ *cobra.Command, _ []string) error {
	a.Request()

	hasFiles := a.ExistFiles()
	hasErrors := a.ExistErrors()
	showErrors := a.ShowErrors()

	if hasFiles {
		if err := a.PrintFiles(); err != nil {
			return err
		}
	}

	if !hasFiles && !hasErrors {
		return a.PrintMessage("No files found.")
	}

	if hasErrors {
		if showErrors {
			return a.PrintErrors()
		}
		return a.PrintMessage("Failed to delete one or more files from the OpenAI API.")
	}

	return nil
}

func (a *DeleteAdapter) Request() {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().FileService()

	a.ListAdapter.Request()
	for _, f := range a.Files() {
		req := &ctrsvc.DeleteFileRequest{FileID: f.ID}
		resp, err := svc.DeleteFile(ctx, req)

		if err != nil {
			a.AddError(err)
		}
		f.ExecStatus = resp.Deleted
	}
}

func (a *DeleteAdapter) PrintFiles() error {
	_ = a.CreateTable()

	a.AppendTableHeader("File ID", "File Name", "Purpose", "Size", "Created", "State")
	a.SetColumnTableConfigs(
		table.ColumnConfig{Number: 1, Align: text.AlignCenter, WidthMin: 27, Colors: text.Colors{text.Bold}},
		table.ColumnConfig{Number: 2, Align: text.AlignRight, WidthMin: 19},
		table.ColumnConfig{Number: 3, Align: text.AlignRight, WidthMin: 19},
		table.ColumnConfig{Number: 4, Align: text.AlignRight, WidthMin: 10},
		table.ColumnConfig{Number: 5, Align: text.AlignRight, WidthMin: 19},
		table.ColumnConfig{Number: 6, Align: text.AlignCenter, WidthMin: 7, Colors: text.Colors{text.Bold}},
	)

	doth := inthlp.EmptyTableColumn
	for _, file := range a.Files() {
		a.AppendTableRow(
			a.FormatString(file.ID, &doth),
			a.FormatString(file.Filename, &doth),
			a.FormatString(file.Purpose, &doth),
			a.FormatBytes(file.Bytes, &doth),
			a.FormatTime(file.CreatedAt, &doth),
			a.FormatExecStatus(file.ExecStatus),
		)
	}

	a.RenderTable()
	return nil
}
