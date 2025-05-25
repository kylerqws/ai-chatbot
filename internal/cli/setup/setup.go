package setup

import (
	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
	"github.com/spf13/cobra"
)

func GeneralConfigure(adp ctr.GeneralAdapter) *cobra.Command {
	cmd := adp.Command()

	PrepareUseField(cmd)
	DisableSortingFlags(cmd)

	FixHelpFunc(cmd)
	HideHelpCommand(cmd)

	return cmd
}

func ParentConfigure(adp ctr.ParentAdapter) *cobra.Command {
	cmd := adp.Command()
	AddCommands(cmd, adp.Children()...)

	return GeneralConfigure(adp)
}

func RootConfigure(adp ctr.RootAdapter) *cobra.Command {
	cmd := adp.Command()
	FixVersionTemplate(cmd, adp.App().Name())

	HideCompletionCommand(cmd)
	SilenceCommandOutput(cmd)

	AddHelpFlag(cmd)
	AddVersionFlag(cmd)

	return ParentConfigure(adp)
}

func CommandConfigure(adp ctr.CommandAdapter) *cobra.Command {
	cmd := adp.Command()
	AddErrorFlag(cmd)

	return GeneralConfigure(adp)
}
