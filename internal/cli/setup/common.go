package setup

import (
	"fmt"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command, children ...*cobra.Command) {
	cmd.AddCommand(children...)
}

func HideHelpCommand(cmd *cobra.Command) {
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
}

func HideCompletionCommand(cmd *cobra.Command) {
	cmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd: true,
	}
}

func SilenceCommandOutput(cmd *cobra.Command) {
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
}

func FixVersionTemplate(cmd *cobra.Command, name string) {
	cmd.SetVersionTemplate(fmt.Sprintf("%s version {{.Version}}\n", name))
}

func FixHelpFunc(cmd *cobra.Command) {
	cmd.SetHelpFunc(HelpFunction())
}

func AddHelpFlag(cmd *cobra.Command) {
	AddBoolFlag(cmd, "help", "h", false, "Show help information", true)
}

func AddVersionFlag(cmd *cobra.Command) {
	AddBoolFlag(cmd, "version", "v", false, "Show application version", false)
}

func AddErrorFlag(cmd *cobra.Command) {
	AddBoolFlag(cmd, "error", "e", false, "Show execution errors", true)
}

func DisableSortingFlags(cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
	cmd.PersistentFlags().SortFlags = false
	cmd.InheritedFlags().SortFlags = false
}
