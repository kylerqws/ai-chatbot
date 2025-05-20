package file

import (
	"fmt"
	"path/filepath"
	"strings"

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
	defaultPurposeFlagKey = "default-purpose"
)

var (
	PurposeManager = enmset.NewPurposeManager()
)

type UploadAdapter struct {
	*inthlp.CommandAdapterHelper
	*inthlp.FlagAdapterHelper
	*inthlp.PrintAdapterHelper
	*inthlp.TableAdapterHelper
	*inthlp.OpenAiFileAdapterHelper
}

func NewUploadAdapter(app *intapp.App) ctr.CommandAdapter {
	adp := &UploadAdapter{}
	cmd := &cobra.Command{}

	adp.CommandAdapterHelper = inthlp.NewCommandAdapterHelper(adp, app, cmd)
	adp.FlagAdapterHelper = inthlp.NewFlagAdapterHelper(cmd)
	adp.PrintAdapterHelper = inthlp.NewPrintAdapterHelper(cmd)
	adp.TableAdapterHelper = inthlp.NewTableAdapterHelper(cmd)
	adp.OpenAiFileAdapterHelper = inthlp.NewOpenAiFileAdapterHelper(cmd)

	return adp
}

func (a *UploadAdapter) Configure() *cobra.Command {
	a.SetUse("upload <file-path[:purpose]> [file-path[:purpose]...]")
	a.SetShort("Upload one or more files to OpenAI")

	a.SetFuncArgs(a.FuncArgs)
	a.SetFuncRunE(a.FuncRunE)

	a.AddFlags()
	return a.MainConfigure()
}

func (a *UploadAdapter) AddFlags() {
	var desc string

	desc = "Default purpose for files without :purpose"
	a.AddStringFlag(defaultPurposeFlagKey, "", "", desc)
}

func (a *UploadAdapter) FuncArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("at least one file path must be specified, usage: %s", cmd.Use)
	}

	if !a.ValidatePurposes() {
		var buf strings.Builder
		if a.ShowErrors() {
			if err := a.PrintErrorsToWriter(&buf); err != nil {
				return err
			}
		}

		return fmt.Errorf("one or more purpose values are incorrect, usage: %s%s",
			PurposeManager.JoinCodes(", "), strings.TrimRight("\n"+buf.String(), "\n"))
	}

	return nil
}

func (a *UploadAdapter) FuncRunE(_ *cobra.Command, _ []string) error {
	a.Request()

	hasFiles := a.ExistFiles()
	hasErrors := a.ExistErrors()
	showErrors := a.ShowErrors()

	if hasFiles {
		if err := a.PrintFiles(); err != nil {
			return err
		}
	}

	if hasErrors {
		if showErrors {
			return a.PrintErrors()
		}
		return a.PrintMessage("Failed to upload one or more files to the OpenAI API.")
	}

	return nil
}

func (a *UploadAdapter) Request() {
	app := a.App()
	ctx := app.Context()
	svc := app.OpenAI().FileService()
	fgs := a.Command().Flags()

	defaultPrpCode, _ := fgs.GetString(defaultPurposeFlagKey)

	for _, arg := range fgs.Args() {
		filePath, prpCode := a.SeparateArg(arg)
		if prpCode == "" {
			prpCode = defaultPrpCode
		}

		prp, _ := PurposeManager.Resolve(prpCode)

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

	a.AppendTableHeader("File ID", "File Name", "Purpose", "Size", "Uploaded")
	a.SetColumnTableConfigs(
		table.ColumnConfig{Number: 1, Align: text.AlignLeft, WidthMin: 27},
		table.ColumnConfig{Number: 2, Align: text.AlignRight, WidthMin: 19},
		table.ColumnConfig{Number: 3, Align: text.AlignRight, WidthMin: 19},
		table.ColumnConfig{Number: 4, Align: text.AlignRight, WidthMin: 10},
		table.ColumnConfig{Number: 5, Align: text.AlignRight, WidthMin: 8},
	)

	doth := inthlp.EmptyTableColumn
	for _, file := range a.Files() {
		a.AppendTableRow(file.ID, file.Filename, file.Purpose,
			a.FormatBytes(file.Bytes, &doth), a.FormatExecStatus(file.ExecStatus),
		)
	}

	a.RenderTable()
	return nil
}

func (_ *UploadAdapter) SeparateArg(arg string) (string, string) {
	parts := strings.SplitN(arg, ":", 2)
	if len(parts) == 2 && parts[1] != "" {
		return parts[0], parts[1]
	}
	return parts[0], ""
}

func (a *UploadAdapter) ValidatePurposes() bool {
	ok := true

	fgs := a.Command().Flags()
	args := fgs.Args()

	for _, arg := range args {
		_, prpCode := a.SeparateArg(arg)
		if _, err := PurposeManager.Resolve(prpCode); err != nil {
			a.AddError(fmt.Errorf("invalid purpose in %q: %w", arg, err))
			ok = false
		}
	}

	if prpCode, err := fgs.GetString(defaultPurposeFlagKey); err != nil {
		a.AddError(fmt.Errorf("failed to get --%s flag value: %w", defaultPurposeFlagKey, err))
		ok = false
	} else if _, err = PurposeManager.Resolve(prpCode); err != nil {
		a.AddError(fmt.Errorf("invalid purpose in --%s flag: %w", defaultPurposeFlagKey, err))
		ok = false
	}

	return ok
}
