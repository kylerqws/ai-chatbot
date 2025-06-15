package setup

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	ctr "github.com/kylerqws/chatbot/internal/cli/contract"
)

const (
	defaultFlagSuffix    = "[flags]"
	customFlagSuffix     = "[flag...]"
	defaultCommandSuffix = "[command]"
	deprecatedMarker     = "[DEPRECATED]"
)

// HelpFunction returns a customized help function for Cobra commands.
func HelpFunction() ctr.FuncHelp {
	return func(cmd *cobra.Command, _ []string) {
		if cmd.Deprecated != "" {
			return
		}

		w := cmd.OutOrStdout()
		sub, loc, glob := cmd.Commands(), localFlags(cmd), globalFlags(cmd)
		hasLoc, hasGlob := existFlags(loc), existFlags(glob)
		hasCmds, hasFlags := existCommands(sub), hasLoc || hasGlob

		if cmd.Short != "" {
			if _, err := fmt.Fprintln(w, cmd.Short); err != nil {
				return
			}
			if _, err := fmt.Fprintln(w); err != nil {
				return
			}
		}
		if cmd.Long != "" {
			if _, err := fmt.Fprintln(w, cmd.Long); err != nil {
				return
			}
			if _, err := fmt.Fprintln(w); err != nil {
				return
			}
		}

		if _, err := fmt.Fprintln(w, "Usage:"); err != nil {
			return
		}
		if _, err := fmt.Fprintf(w, "  %s\n", useLine(cmd, hasCmds, hasFlags)); err != nil {
			return
		}

		if hasCmds {
			if _, err := fmt.Fprintln(w, "\nCommands:"); err != nil {
				return
			}
			for i := range sub {
				if err := printCommandLine(w, sub[i]); err != nil {
					return
				}
			}
		}

		if hasFlags {
			if _, err := fmt.Fprintln(w, "\nFlags:"); err != nil {
				return
			}
			for i := range loc {
				if err := printFlagLine(w, loc[i]); err != nil {
					return
				}
			}
			if hasLoc && hasGlob && len(glob) > 1 {
				if _, err := fmt.Fprintln(w); err != nil {
					return
				}
			}
			for i := range glob {
				if err := printFlagLine(w, glob[i]); err != nil {
					return
				}
			}
		}
	}
}

// useLine generates the usage line.
func useLine(cmd *cobra.Command, existCmds, existFlags bool) string {
	line := strings.TrimSpace(strings.TrimSuffix(cmd.UseLine(), defaultFlagSuffix))
	if existCmds {
		line += " " + defaultCommandSuffix
	}
	if existFlags {
		line += " " + customFlagSuffix
	}
	return line
}

// localFlags returns local flags that are not inherited or persistent.
func localFlags(cmd *cobra.Command) []*pflag.Flag {
	var flags []*pflag.Flag
	i, p := cmd.InheritedFlags(), cmd.PersistentFlags()

	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		if !f.Hidden && i.Lookup(f.Name) == nil && p.Lookup(f.Name) == nil {
			flags = append(flags, f)
		}
	})

	return flags
}

// globalFlags returns all inherited and persistent flags.
func globalFlags(cmd *cobra.Command) []*pflag.Flag {
	var flags []*pflag.Flag

	cmd.InheritedFlags().VisitAll(func(f *pflag.Flag) {
		if !f.Hidden {
			flags = append(flags, f)
		}
	})

	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		if !f.Hidden {
			flags = append(flags, f)
		}
	})

	return flags
}

// existCommands reports whether any visible subcommands are present.
func existCommands(cmds []*cobra.Command) bool {
	for i := range cmds {
		if !cmds[i].Hidden {
			return true
		}
	}
	return false
}

// existFlags reports whether any visible flags are present.
func existFlags(flags []*pflag.Flag) bool {
	for i := range flags {
		if !flags[i].Hidden {
			return true
		}
	}
	return false
}

// printCommandLine writes a formatted subcommand entry.
func printCommandLine(w io.Writer, cmd *cobra.Command) error {
	if cmd.Hidden {
		return nil
	}

	name, desc := fmt.Sprintf("  %s", cmd.Name()), cmd.Short
	if cmd.Deprecated != "" {
		desc += " " + deprecatedMarker
	}

	if _, err := fmt.Fprintf(w, "%-20s\t%s\n", name, desc); err != nil {
		return err
	}
	return nil
}

// printFlagLine writes a formatted flag entry.
func printFlagLine(w io.Writer, flag *pflag.Flag) error {
	if flag.Hidden {
		return nil
	}

	var name string
	if flag.Shorthand != "" {
		name = fmt.Sprintf("  -%s, --%s", flag.Shorthand, flag.Name)
	} else {
		name = fmt.Sprintf("  --%s", flag.Name)
	}

	if _, err := fmt.Fprintf(w, "%-20s\t%s\n", name, flag.Usage); err != nil {
		return err
	}
	return nil
}
