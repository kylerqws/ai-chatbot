package file

import (
	"path/filepath"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	hlpfil "github.com/kylerqws/chatbot/internal/cli/helper/adapter/cmd/openai/file"
	hlpcmd "github.com/kylerqws/chatbot/internal/cli/helper/adapter/command"
	hlpfmt "github.com/kylerqws/chatbot/internal/cli/helper/adapter/format"
	hlptbl "github.com/kylerqws/chatbot/internal/cli/helper/adapter/table"

	ctradp "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

const (
	defaultPurposeFlagKey   = "default-purpose"
	argumentSeparatorSymbol = ":"
)

type UploadAdapter struct {
	*hlpcmd.CommandAdapterHelper
	*hlptbl.TableAdapterHelper
	*hlpfmt.FormatAdapterHelper
	*hlpfil.ValidateOpenAiFileAdapterHelper
}

func NewUploadAdapter(app *intapp.App) ctradp.CommandAdapter {
	adp := &UploadAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapterHelper = hlpcmd.NewCommandAdapterHelper(app, cmd)
	adp.TableAdapterHelper = hlptbl.NewTableAdapterHelper(cmd)
	adp.FormatAdapterHelper = hlpfmt.NewFormatAdapterHelper(cmd)
	adp.ValidateOpenAiFileAdapterHelper = hlpfil.NewValidateOpenAiFileAdapterHelper(cmd)

	return adp
}

func (a *UploadAdapter) Configure() *cobra.Command {
	a.SetUse("upload <file-path[:purpose]> [file-path[:purpose]...]")
	a.SetShort("Upload one or more files to OpenAI")

	a.SetFuncArgs(a.Validate)
	a.SetFuncRunE(a.Upload)

	a.ConfigureFlags()
	return a.MainConfigure()
}

func (a *UploadAdapter) ConfigureFlags() {
	desc := "Default purpose for files without :purpose suffix"
	a.AddStringFlag(defaultPurposeFlagKey, "", "", desc)
}

func (a *UploadAdapter) Validate(cmd *cobra.Command, args []string) error {
	a.AddError(a.ValidateHasAnyArgsMore(0))
	a.AddErrors(a.ValidatePurposes()...)

	return a.ErrorIfExist("one or more arguments are invalid or missing")
}

func (a *UploadAdapter) ValidatePurposes() []error {
	var errs []error
	args := a.Command().Flags().Args()

	for i := range args {
		_, prpCode := a.SeparateArg(args[i])
		if err := a.ValidatePurposeCode(prpCode); err != nil {
			errs = append(errs, err)
		}
	}

	if err := a.ValidatePurposeFlag(defaultPurposeFlagKey); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func (a *UploadAdapter) Upload(_ *cobra.Command, _ []string) error {
	a.Request()

	if a.ExistFiles() {
		if err := a.PrintFiles(); err != nil {
			return err
		}
	}

	return a.ErrorIfExist("failed to upload files or data is unavailable")
}

func (a *UploadAdapter) Request() {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().FileService()
	mng := a.PurposeManager()
	fgs := a.Command().Flags()
	args := fgs.Args()

	defaultPrpCode, err := fgs.GetString(defaultPurposeFlagKey)
	if err != nil {
		a.AddError(err)
	}

	for i := range args {
		filePath, prpCode := a.SeparateArg(args[i])
		if prpCode == "" {
			prpCode = defaultPrpCode
		}

		prp, _ := mng.Resolve(prpCode)
		resp, err := svc.UploadFile(ctx, &ctrsvc.UploadFileRequest{
			FilePath: filePath,
			Purpose:  prp.Code,
		})

		if err != nil {
			a.AddError(err)

			resp.File = &ctrsvc.File{
				Filename: filepath.Base(filePath),
				Purpose:  prp.Code,
				Bytes:    a.FileSize(filePath),
			}
		}

		file := a.AddFile(resp.File)
		if err == nil {
			file.ExecStatus = true
		}
	}
}

func (a *UploadAdapter) PrintFiles() error {
	_ = a.CreateTable()

	a.AppendTableHeader("File ID", "File Name", "Purpose", "Size", "Created", "State")
	a.SetColumnTableConfigs(
		a.ColumnConfig(1, text.AlignCenter, 27, text.Colors{text.Bold}),
		a.ColumnConfig(2, text.AlignRight, 19, nil),
		a.ColumnConfig(3, text.AlignRight, 19, nil),
		a.ColumnConfig(4, text.AlignRight, 10, nil),
		a.ColumnConfig(5, text.AlignRight, 19, nil),
		a.ColumnConfig(6, text.AlignCenter, 7, text.Colors{text.Bold}),
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
			a.FormatExecStatus(files[i].ExecStatus),
		)
	}

	a.RenderTable()
	return nil
}

func (*UploadAdapter) SeparateArg(arg string) (string, string) {
	parts := strings.SplitN(arg, argumentSeparatorSymbol, 2)
	if len(parts) == 2 && parts[1] != "" {
		return parts[0], parts[1]
	}
	return parts[0], ""
}
