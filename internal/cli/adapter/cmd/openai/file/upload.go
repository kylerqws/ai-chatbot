package file

import (
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	intapp "github.com/kylerqws/chatbot/internal/app"
	helper "github.com/kylerqws/chatbot/internal/cli/helper/adapter"

	ctradp "github.com/kylerqws/chatbot/internal/cli/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type UploadAdapter struct {
	*helper.CommandAdapter
	*helper.OpenAiAdapter
	*helper.OpenAiFileAdapter
	*helper.ArgumentAdapter
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
	adp.ArgumentAdapter = helper.NewArgumentAdapter(cmd)
	adp.FlagAdapter = helper.NewFlagAdapter(cmd)
	adp.ValidateAdapter = helper.NewValidateAdapter(cmd)
	adp.TableAdapter = helper.NewTableAdapter(cmd)
	adp.FormatAdapter = helper.NewFormatAdapter(cmd)

	return adp
}

func (a *UploadAdapter) Configure() *cobra.Command {
	argExample := "file-path[:purpose]"
	a.SetUse("upload <" + argExample + "> [" + argExample + "...]")
	a.SetShort("Upload one or more files to OpenAI account")
	a.SetLong(a.FileUploadHelpInfo(a.OpenAiAdapter))

	a.SetFuncArgs(a.Validate)
	a.SetFuncRunE(a.Upload)

	a.ConfigureFlags()
	return a.MainConfigure()
}

func (a *UploadAdapter) ConfigureFlags() {
	prpManager := a.PurposeManager()

	desc := "Default purpose for files without ':purpose' suffix\n" +
		"\t\t\t(" + prpManager.JoinCodes(", ") + ")"
	a.AddStringFlag(helper.DefaultPurposeFlagKey, "", prpManager.Default().Code, desc)
}

func (a *UploadAdapter) Validate(_ *cobra.Command, _ []string) error {
	a.AddError(a.ValidateHasMoreArgsThan(0))

	a.AddErrors(a.ValidateArgs()...)
	a.AddErrors(a.ValidateFlags()...)

	return a.ErrorIfExist("One or more arguments/flags are invalid or missing.")
}

func (a *UploadAdapter) ValidateArgs() []error {
	var errs []error
	args := a.Command().Flags().Args()

	for i := range args {
		_, prpCode := a.SplitArg(args[i])

		err := a.ValidatePurposeCode(prpCode)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (a *UploadAdapter) ValidateFlags() []error {
	var errs []error

	err := a.ValidateStringFlag(helper.DefaultPurposeFlagKey, a.ValidatePurposeCode)
	if err != nil {
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

	return a.ErrorIfExist("Failed to upload files or data is unavailable.")
}

func (a *UploadAdapter) Request() bool {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().FileService()
	fgs := a.Command().Flags()
	args := fgs.Args()

	defaultPrpCode, err := fgs.GetString(helper.DefaultPurposeFlagKey)
	if err != nil {
		a.AddError(err)
	}

	for i := range args {
		filePath, prpCode := a.SplitArg(args[i], defaultPrpCode)

		resp, err := svc.UploadFile(ctx, &ctrsvc.UploadFileRequest{
			FilePath: filePath,
			Purpose:  prpCode,
		})

		if err != nil {
			a.AddError(err)

			resp.File = &ctrsvc.File{
				Filename: a.FileName(filePath),
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

	a.AppendTableHeader("File ID", "File Name", "Purpose",
		"Size", "Created", "Status", "State")

	a.SetColumnTableConfigs(
		a.ColumnConfig(1, text.AlignCenter, 27, text.Colors{text.Bold}),
		a.ColumnConfig(2, text.AlignRight, 19, nil),
		a.ColumnConfig(3, text.AlignRight, 19, nil),
		a.ColumnConfig(4, text.AlignRight, 10, nil),
		a.ColumnConfig(5, text.AlignRight, 19, nil),
		a.ColumnConfig(6, text.AlignRight, 10, nil),
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
			a.FormatTime(files[i].CreatedAt, &empty),
			a.FormatString(files[i].Status, &empty),
			a.FormatExecStatus(files[i].ExecStatus),
		)
	}

	a.RenderTable()
	return nil
}

func (a *UploadAdapter) SplitArg(arg string, defaults ...string) (string, string) {
	filePath, prpCode := a.SplitDoubleArg(arg)
	defCount := len(defaults)

	if prpCode == "" && defCount > 0 {
		prpCode = defaults[0]
	}

	return filePath, prpCode
}
