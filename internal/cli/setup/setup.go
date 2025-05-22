package setup

import (
	"github.com/spf13/cobra"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
	cmn "github.com/kylerqws/chatbot/internal/cli/setup/common"
)

func GeneralConfigure(adp ctr.Adapter) *cobra.Command {
	cmd := adp.Command()

	cmn.PrepareUseField(cmd)
	cmn.DisableSortingFlag(cmd)

	cmn.FixHelpFunc(cmd)
	cmn.HideHelpCommand(cmd)

	return cmd
}

func RootConfigure(adp ctr.RootAdapter) *cobra.Command {
	cmd := adp.Command()

	cmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd: true,
	}

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	cmn.FixVersionTemplate(cmd, adp.App().Name())

	cmn.AddHelpFlag(cmd)
	cmn.AddVersionFlag(cmd)

	return ParentConfigure(adp)
}

func ParentConfigure(adp ctr.ParentAdapter) *cobra.Command {
	cmd := adp.Command()
	cmn.AddCommands(cmd, adp.Children()...)

	return GeneralConfigure(adp)
}

func CommandConfigure(adp ctr.CommandAdapter) *cobra.Command {
	cmd := adp.Command()
	cmn.AddErrorFlag(cmd)

	return GeneralConfigure(adp)
}
