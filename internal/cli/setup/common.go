package setup

import (
	"fmt"
	"github.com/spf13/cobra"
)

// AddCommands registers the given subcommands under the specified command.
func AddCommands(cmd *cobra.Command, children ...*cobra.Command) {
	cmd.AddCommand(children...)
}

// HideHelpCommand disables the built-in help command.
func HideHelpCommand(cmd *cobra.Command) {
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
}

// HideCompletionCommand disables the built-in completion command.
func HideCompletionCommand(cmd *cobra.Command) {
	cmd.CompletionOptions.DisableDefaultCmd = true
}

// SilenceCommandOutput suppresses usage and error output from the command.
func SilenceCommandOutput(cmd *cobra.Command) {
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
}

// FixVersionTemplate sets a custom version output format.
func FixVersionTemplate(cmd *cobra.Command, name string) {
	cmd.SetVersionTemplate(fmt.Sprintf("%s version {{.Version}}\n", name))
}

// FixHelpFunc sets a custom help function for the command.
func FixHelpFunc(cmd *cobra.Command) {
	cmd.SetHelpFunc(HelpFunction())
}

// AddHelpFlag registers the `--help` flag.
func AddHelpFlag(cmd *cobra.Command) {
	AddBoolFlag(cmd, "help", "h", false, "Show help information", true)
}

// AddVersionFlag registers the `--version` flag.
func AddVersionFlag(cmd *cobra.Command) {
	AddBoolFlag(cmd, "version", "v", false, "Show application version", false)
}

// AddErrorFlag registers the `--error` flag.
func AddErrorFlag(cmd *cobra.Command) {
	AddBoolFlag(cmd, "error", "e", false, "Show execution errors", true)
}

// DisableSortingFlags disables automatic sorting of flags.
func DisableSortingFlags(cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
	cmd.PersistentFlags().SortFlags = false
	cmd.InheritedFlags().SortFlags = false
}
