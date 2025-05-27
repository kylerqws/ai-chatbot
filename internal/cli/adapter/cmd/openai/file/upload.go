package file

import (
	"path/filepath"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctradp "github.com/kylerqws/chatbot/internal/cli/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

const (
	defaultPurposeFlagKey   = "default-purpose"
	argumentSeparatorSymbol = ":"
)

type UploadAdapter struct {
	*helper.CommandAdapter
	*helper.OpenAiAdapter
	*helper.OpenAiFileAdapter
	*helper.FlagAdapter
	*helper.ValidateAdapter
	*helper.TableAdapter
	*helper.FormatAdapter
}

func NewUploadAdapter(app *intapp.App) ctradp.CommandAdapter {
	adp := &UploadAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapter = helper.NewCommandAdapter(app, cmd)
	adp.OpenAiAdapter = helper.NewOpenAiAdapter(cmd)
	adp.OpenAiFileAdapter = helper.NewOpenAiFileAdapter(cmd)
	adp.FlagAdapter = helper.NewFlagAdapter(cmd)
	adp.ValidateAdapter = helper.NewValidateAdapter(cmd)
	adp.TableAdapter = helper.NewTableAdapter(cmd)
	adp.FormatAdapter = helper.NewFormatAdapter(cmd)

	return adp
}

func (a *UploadAdapter) Configure() *cobra.Command {
	a.SetUse("upload <file-path[:purpose]> [file-path[:purpose]...]")
	a.SetShort("Upload one or more files to OpenAI account")

	a.SetFuncArgs(a.Validate)
	a.SetFuncRunE(a.Upload)

	a.ConfigureFlags()
	return a.MainConfigure()
}

func (a *UploadAdapter) ConfigureFlags() {
	desc := "Default purpose for files without ':purpose' suffix"
	a.AddStringFlag(defaultPurposeFlagKey, "", "", desc)
}

func (a *UploadAdapter) Validate(_ *cobra.Command, _ []string) error {
	a.AddError(a.ValidateHasMoreArgsThan(0))
	a.AddErrors(a.ValidatePurposeCodes()...)

	return a.ErrorIfExist("one or more arguments are invalid or missing")
}

func (a *UploadAdapter) ValidatePurposeCodes() []error {
	var errs []error
	args := a.Command().Flags().Args()

	for i := range args {
		_, prpCode := a.separateArg(args[i])
		if err := a.ValidatePurposeCode(prpCode); err != nil {
			errs = append(errs, err)
		}
	}

	if err := a.ValidateStringFlag(defaultPurposeFlagKey, a.ValidatePurposeCode); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func (a *UploadAdapter) Upload(_ *cobra.Command, _ []string) error {
	if a.Request() && a.ExistFiles() {
		if err := a.PrintFiles(); err != nil {
			return err
		}
	}

	return a.ErrorIfExist("failed to upload files or data is unavailable")
}

func (a *UploadAdapter) Request() bool {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().FileService()
	fgs := a.Command().Flags()
	args := fgs.Args()

	defaultPrpCode, err := fgs.GetString(defaultPurposeFlagKey)
	if err != nil {
		a.AddError(err)
	}

	for i := range args {
		filePath, prpCode := a.separateArg(args[i])
		if prpCode == "" {
			prpCode = defaultPrpCode
		}

		resp, err := svc.UploadFile(ctx, &ctrsvc.UploadFileRequest{
			FilePath: filePath,
			Purpose:  prpCode,
		})

		if err != nil {
			a.AddError(err)

			resp.File = &ctrsvc.File{
				Filename: filepath.Base(filePath),
				Purpose:  prpCode,
				Bytes:    a.FileSize(filePath),
			}
		}

		file := a.WrapOpenAIFile(resp.File)
		file.ExecStatus = err == nil

		a.AddFile(file)
	}

	return true
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
	doth := helper.EmptyTableColumn

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

func (*UploadAdapter) separateArg(arg string) (string, string) {
	parts := strings.SplitN(arg, argumentSeparatorSymbol, 2)
	if len(parts) == 2 && parts[1] != "" {
		return parts[0], parts[1]
	}
	return parts[0], ""
}
