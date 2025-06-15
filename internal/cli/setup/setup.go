package setup

import (
	ctr "github.com/kylerqws/chatbot/internal/cli/contract/adapter"
	"github.com/spf13/cobra"
)

// GeneralConfigure applies common configuration for an all CLI command.
func GeneralConfigure(adp ctr.GeneralAdapter) *cobra.Command {
	cmd := adp.Command()
	DisableSortingFlags(cmd)

	FixHelpFunc(cmd)
	HideHelpCommand(cmd)

	return cmd
}

// ParentConfigure applies common configuration for a parent CLI command with subcommands.
func ParentConfigure(adp ctr.ParentAdapter) *cobra.Command {
	cmd := adp.Command()
	AddCommands(cmd, adp.Children()...)

	return GeneralConfigure(adp)
}

// RootConfigure applies configuration for the root CLI command.
func RootConfigure(adp ctr.RootAdapter) *cobra.Command {
	cmd := adp.Command()
	FixVersionTemplate(cmd, adp.App().Name())

	HideCompletionCommand(cmd)
	SilenceCommandOutput(cmd)

	AddHelpFlag(cmd)
	AddVersionFlag(cmd)

	return ParentConfigure(adp)
}

// CommandConfigure applies configuration for a functional CLI command.
func CommandConfigure(adp ctr.CommandAdapter) *cobra.Command {
	cmd := adp.Command()
	AddErrorFlag(cmd)

	return GeneralConfigure(adp)
}
