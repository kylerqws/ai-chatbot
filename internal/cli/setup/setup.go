package setup

import (
	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
	"github.com/spf13/cobra"
)

// GeneralConfigure applies standard configuration to all commands.
func GeneralConfigure(adp ctr.GeneralAdapter) *cobra.Command {
	cmd := adp.Command()
	DisableSortingFlags(cmd)

	FixHelpFunc(cmd)
	HideHelpCommand(cmd)

	return cmd
}

// ParentConfigure applies standard configuration to a parent command with subcommands.
func ParentConfigure(adp ctr.ParentAdapter) *cobra.Command {
	cmd := adp.Command()
	AddCommands(cmd, adp.Children()...)

	return GeneralConfigure(adp)
}

// RootConfigure applies configuration specific to the root command.
func RootConfigure(adp ctr.RootAdapter) *cobra.Command {
	cmd := adp.Command()
	FixVersionTemplate(cmd, adp.App().Name())

	HideCompletionCommand(cmd)
	SilenceCommandOutput(cmd)

	AddHelpFlag(cmd)
	AddVersionFlag(cmd)

	return ParentConfigure(adp)
}

// CommandConfigure applies configuration to a functional command with logic and error output.
func CommandConfigure(adp ctr.CommandAdapter) *cobra.Command {
	cmd := adp.Command()
	AddErrorFlag(cmd)

	return GeneralConfigure(adp)
}
