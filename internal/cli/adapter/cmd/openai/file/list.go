package file

import (
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	hlpfil "github.com/kylerqws/chatbot/internal/cli/helper/adapter/cmd/openai/file"
	hlpcmd "github.com/kylerqws/chatbot/internal/cli/helper/adapter/command"
	hlpdmt "github.com/kylerqws/chatbot/internal/cli/helper/adapter/datetime"
	hlpfmt "github.com/kylerqws/chatbot/internal/cli/helper/adapter/format"
	hlptbl "github.com/kylerqws/chatbot/internal/cli/helper/adapter/table"

	ctradp "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

const (
	idFlagKey            = "id"
	purposeFlagKey       = "purpose"
	createdAfterFlagKey  = "created-after"
	createdBeforeFlagKey = "created-before"
	fileIDExample        = "file-xxxxxx..."
	dateExample          = "1970-01-01"
	datetimeExample      = "1970-01-01 00:00:00"
)

type ListAdapter struct {
	*hlpcmd.CommandAdapterHelper
	*hlptbl.TableAdapterHelper
	*hlpfmt.FormatAdapterHelper
	*hlpdmt.ValidateDateTimeAdapterHelper
	*hlpfil.ValidateOpenAiFileAdapterHelper
}

func NewListAdapter(app *intapp.App) ctradp.CommandAdapter {
	adp := &ListAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapterHelper = hlpcmd.NewCommandAdapterHelper(app, cmd)
	adp.TableAdapterHelper = hlptbl.NewTableAdapterHelper(cmd)
	adp.FormatAdapterHelper = hlpfmt.NewFormatAdapterHelper(cmd)
	adp.ValidateDateTimeAdapterHelper = hlpdmt.NewValidateDateTimeAdapterHelper(cmd)
	adp.ValidateOpenAiFileAdapterHelper = hlpfil.NewValidateOpenAiFileAdapterHelper(cmd)

	return adp
}

func (a *ListAdapter) Configure() *cobra.Command {
	a.SetUse("list")
	a.SetShort("List files in OpenAI account")

	a.SetFuncArgs(a.Validate)
	a.SetFuncRunE(a.List)

	a.ConfigureFlags()
	return a.MainConfigure()
}

func (a *ListAdapter) ConfigureFlags() {
	desc := "Filter by file ID (e.g. " + fileIDExample + ")"
	a.AddStringSliceFlag(idFlagKey, "", []string{}, desc)

	desc = "Filter by purpose (e.g. " + a.PurposeManager().JoinCodes(", ") + ")"
	a.AddStringSliceFlag(purposeFlagKey, "", []string{}, desc)

	desc = "Filter by creation date after (e.g. " + dateExample + " or " + datetimeExample + ")"
	a.AddStringFlag(createdAfterFlagKey, "", a.DateTime(0, -1, 0), desc)

	desc = "Filter by creation date before (e.g. " + dateExample + " or " + datetimeExample + ")"
	a.AddStringFlag(createdBeforeFlagKey, "", "", desc)
}

func (a *ListAdapter) Validate(_ *cobra.Command, _ []string) error {
	a.AddErrors(a.ValidateFileIDFlags(idFlagKey)...)
	a.AddErrors(a.ValidatePurposeFlags(purposeFlagKey)...)

	a.AddError(a.ValidateDateFlag(createdAfterFlagKey))
	a.AddError(a.ValidateDateFlag(createdBeforeFlagKey))

	return a.ErrorIfExist("one or more arguments are invalid or missing")
}

func (a *ListAdapter) List(_ *cobra.Command, _ []string) error {
	a.Request()

	if a.ExistFiles() {
		if err := a.PrintFiles(); err != nil {
			return err
		}
	} else if !a.ExistErrors() {
		return a.PrintMessage("No files found.")
	}

	return a.ErrorIfExist("failed to retrieve files or data is unavailable")
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

	purposes, err := fgs.GetStringSlice(purposeFlagKey)
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
		Purposes:      purposes,
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
		a.ColumnConfig(1, text.AlignCenter, 27, text.Colors{text.Bold}),
		a.ColumnConfig(2, text.AlignRight, 19, nil),
		a.ColumnConfig(3, text.AlignRight, 19, nil),
		a.ColumnConfig(4, text.AlignRight, 10, nil),
		a.ColumnConfig(5, text.AlignRight, 19, nil),
	)

	files := a.Files()
	doth := hlptbl.EmptyTableColumn

	for i := range files {
		a.AppendTableRow(
			a.FormatString(files[i].ID, &doth),
			a.FormatString(files[i].Filename, &doth),
			a.FormatString(files[i].Purpose, &doth),
			a.FormatBytes(files[i].Bytes, &doth),
			a.FormatTime(files[i].CreatedAt, &doth),
		)
	}

	a.RenderTable()
	return nil
}
