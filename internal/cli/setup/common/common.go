package common

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/cli/setup/flag"
	"github.com/kylerqws/chatbot/internal/cli/setup/help"
)

func AddCommands(cmd *cobra.Command, children ...*cobra.Command) {
	cmd.AddCommand(children...)
}

func HideHelpCommand(cmd *cobra.Command) {
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
}

func FixVersionTemplate(cmd *cobra.Command, name string) {
	cmd.SetVersionTemplate(fmt.Sprintf("%s version {{.Version}}\n", name))
}

func FixHelpFunc(cmd *cobra.Command) {
	cmd.SetHelpFunc(help.FunctionHelp())
}

func AddHelpFlag(cmd *cobra.Command) {
	flag.AddBoolFlag(cmd, "help", "h", false, "Show help information", true)
}

func AddVersionFlag(cmd *cobra.Command) {
	flag.AddBoolFlag(cmd, "version", "v", false, "Show application version", false)
}

func AddErrorFlag(cmd *cobra.Command) {
	flag.AddBoolFlag(cmd, "error", "e", false, "Show execution errors", true)
}

func DisableSortingFlag(cmd *cobra.Command) {
	flag.DisableSorting(cmd)
}

func PrepareUseField(cmd *cobra.Command) {
	if slices.ContainsFunc(cmd.Commands(), func(c *cobra.Command) bool {
		return !c.Hidden
	}) {
		cmd.Use += " [command]"
	}
}
