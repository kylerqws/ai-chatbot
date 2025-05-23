package setup

import (
	"github.com/kylerqws/chatbot/internal/cli/setup/common"
	"github.com/spf13/cobra"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
	ctradp "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
)

func GeneralConfigure(adp ctr.Adapter) *cobra.Command {
	cmd := adp.Command()

	common.PrepareUseField(cmd)
	common.DisableSortingFlag(cmd)

	common.FixHelpFunc(cmd)
	common.HideHelpCommand(cmd)

	return cmd
}

func RootConfigure(adp ctradp.RootAdapter) *cobra.Command {
	cmd := adp.Command()
	common.FixVersionTemplate(cmd, adp.App().Name())

	common.HideCompletionCommand(cmd)
	common.SilenceCommandOutput(cmd)

	common.AddHelpFlag(cmd)
	common.AddVersionFlag(cmd)

	return ParentConfigure(adp)
}

func ParentConfigure(adp ctradp.ParentAdapter) *cobra.Command {
	cmd := adp.Command()
	common.AddCommands(cmd, adp.Children()...)

	return GeneralConfigure(adp)
}

func CommandConfigure(adp ctradp.CommandAdapter) *cobra.Command {
	cmd := adp.Command()
	common.AddErrorFlag(cmd)

	return GeneralConfigure(adp)
}
