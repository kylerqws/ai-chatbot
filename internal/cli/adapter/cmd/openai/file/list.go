package file

import (
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	intapp "github.com/kylerqws/chatbot/internal/app"
	inthlp "github.com/kylerqws/chatbot/internal/cli/helper"
	enmset "github.com/kylerqws/chatbot/internal/openai/enumset"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

const (
	idFlagKey            = "id"
	purposeFlagKey       = "purpose"
	createdAfterFlagKey  = "created-after"
	createdBeforeFlagKey = "created-before"
	fileIDTemplate       = "file-xxxxxx..."
	dateTemplate         = "1970-01-01"
	datetimeTemplate     = "1970-01-01 00:00:00"
)

type ListAdapter struct {
	*inthlp.CommandAdapterHelper
	*inthlp.FlagAdapterHelper
	*inthlp.PrintAdapterHelper
	*inthlp.TableAdapterHelper
	*inthlp.DateTimeAdapterHelper
	*inthlp.OpenAiFileAdapterHelper
}

func NewListAdapter(app *intapp.App) ctr.CommandAdapter {
	adp := &ListAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapterHelper = inthlp.NewCommandAdapterHelper(adp, app, cmd)
	adp.FlagAdapterHelper = inthlp.NewFlagAdapterHelper(cmd)
	adp.PrintAdapterHelper = inthlp.NewPrintAdapterHelper(cmd)
	adp.TableAdapterHelper = inthlp.NewTableAdapterHelper(cmd)
	adp.DateTimeAdapterHelper = inthlp.NewDateTimeAdapterHelper(cmd)
	adp.OpenAiFileAdapterHelper = inthlp.NewOpenAiFileAdapterHelper(cmd)

	return adp
}

func (a *ListAdapter) Configure() *cobra.Command {
	a.SetUse("list")
	a.SetShort("List files in OpenAI account")
	a.SetFuncRunE(a.FuncRunE)

	a.AddFlags()
	return a.MainConfigure()
}

func (a *ListAdapter) AddFlags() {
	var desc string

	desc = "Filter by file ID (e.g. " + fileIDTemplate + ")"
	a.AddStringSliceFlag(idFlagKey, []string{}, desc)

	desc = "Filter by purpose (e.g. " + enmset.NewPurposeManager().JoinCodes(", ") + ")"
	a.AddStringFlag(purposeFlagKey, "", "", desc)

	desc = "Filter by creation date after (e.g. " + dateTemplate + " or " + datetimeTemplate + ")"
	a.AddStringFlag(createdAfterFlagKey, "", a.DateTime(0, -1, 0), desc)

	desc = "Filter by creation date before (e.g. " + dateTemplate + " or " + datetimeTemplate + ")"
	a.AddStringFlag(createdBeforeFlagKey, "", "", desc)
}

func (a *ListAdapter) FuncRunE(_ *cobra.Command, _ []string) error {
	a.Request()
	if a.ExistFiles() {
		return a.PrintFiles()
	}

	if a.ExistErrors() {
		if a.ShowErrors() {
			return a.PrintErrors()
		}
		return a.PrintMessage("Failed to retrieve the file list from the OpenAI API.")
	}

	return a.PrintMessage("No files found.")
}

func (a *ListAdapter) Request() {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().FileService()
	fgs := a.Command().Flags()

	fileIDs, err := fgs.GetStringSlice(idFlagKey)
	if err != nil {
		a.AddError(err)
	}

	purpose, err := fgs.GetString(purposeFlagKey)
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
		Purpose:       purpose,
		CreatedAfter:  a.ParseDateTime(afterStr),
		CreatedBefore: a.ParseDateTime(beforeStr),
	})

	if err != nil {
		a.AddError(err)
	}
	a.AddFiles(resp.Files...)
}

func (a *ListAdapter) PrintFiles() error {
	_ = a.CreateTable()

	a.AppendTableHeader("File ID", "File Name", "Purpose", "Size", "Created")
	a.SetColumnTableConfigs(
		table.ColumnConfig{Number: 1, Align: text.AlignLeft, WidthMin: 27},
		table.ColumnConfig{Number: 2, Align: text.AlignRight, WidthMin: 19},
		table.ColumnConfig{Number: 3, Align: text.AlignRight, WidthMin: 19},
		table.ColumnConfig{Number: 4, Align: text.AlignRight, WidthMin: 10},
		table.ColumnConfig{Number: 5, Align: text.AlignRight, WidthMin: 19},
	)

	doth := inthlp.EmptyTableColumn
	for _, file := range a.Files() {
		a.AppendTableRow(file.ID, file.Filename, file.Purpose,
			a.FormatBytes(file.Bytes, &doth), a.FormatTime(file.CreatedAt, &doth),
		)
	}

	a.RenderTable()
	return nil
}
