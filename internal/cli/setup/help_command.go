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
	oldFlagsSuffix   = "[flags]"
	newFlagsSuffix   = "[flag...]"
	deprecatedSuffix = "[DEPRECATED]"
)

func HelpFunction() ctr.FuncHelp {
	return func(cmd *cobra.Command, _ []string) {
		if cmd.Deprecated != "" {
			return
		}
		w := cmd.OutOrStdout()

		if cmd.Long != "" {
			if _, err := fmt.Fprintln(w, cmd.Long); err != nil {
				return
			}
			if _, err := fmt.Fprintln(w); err != nil {
				return
			}
		} else if cmd.Short != "" {
			if _, err := fmt.Fprintln(w, cmd.Short); err != nil {
				return
			}
			if _, err := fmt.Fprintln(w); err != nil {
				return
			}
		}

		if _, err := fmt.Fprintln(w, "Usage:"); err != nil {
			return
		}
		if _, err := fmt.Fprintf(w, "  %s\n", useLine(cmd)); err != nil {
			return
		}

		sub := cmd.Commands()
		if existCommands(sub) {
			if _, err := fmt.Fprintln(w, "\nCommands:"); err != nil {
				return
			}
			for i := range sub {
				if err := printCommandLine(w, sub[i]); err != nil {
					return
				}
			}
		}

		loc, glob := localFlags(cmd), globalFlags(cmd)
		hasLoc, hasGlob := existFlags(loc), existFlags(glob)

		if hasLoc || hasGlob {
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

func useLine(cmd *cobra.Command) string {
	cmdUseLine := cmd.UseLine()
	if !strings.HasSuffix(cmdUseLine, oldFlagsSuffix) {
		return cmdUseLine
	}

	return strings.TrimSuffix(cmdUseLine, oldFlagsSuffix) + newFlagsSuffix
}

func localFlags(cmd *cobra.Command) []*pflag.Flag {
	var flags []*pflag.Flag

	inheritedFlags := cmd.InheritedFlags()
	persistentFlags := cmd.PersistentFlags()

	if set := cmd.LocalFlags(); set != nil {
		set.VisitAll(func(f *pflag.Flag) {
			if f.Hidden {
				return
			}
			if inheritedFlags.Lookup(f.Name) != nil {
				return
			}
			if persistentFlags.Lookup(f.Name) != nil {
				return
			}

			flags = append(flags, f)
		})
	}

	return flags
}

func globalFlags(cmd *cobra.Command) []*pflag.Flag {
	var flags []*pflag.Flag

	if set := cmd.InheritedFlags(); set != nil {
		set.VisitAll(func(f *pflag.Flag) {
			if f.Hidden {
				return
			}

			flags = append(flags, f)
		})
	}

	if set := cmd.PersistentFlags(); set != nil {
		set.VisitAll(func(f *pflag.Flag) {
			if f.Hidden {
				return
			}

			flags = append(flags, f)
		})
	}

	return flags
}

func existCommands(list []*cobra.Command) bool {
	for i := range list {
		if !list[i].Hidden {
			return true
		}
	}
	return false
}

func existFlags(list []*pflag.Flag) bool {
	for i := range list {
		if !list[i].Hidden {
			return true
		}
	}
	return false
}

func printCommandLine(w io.Writer, cmd *cobra.Command) error {
	if cmd.Hidden {
		return nil
	}

	cmdPart, cmdShort := fmt.Sprintf("  %s", cmd.Name()), cmd.Short
	if cmd.Deprecated != "" {
		cmdShort += " " + deprecatedSuffix
	}

	if _, err := fmt.Fprintf(w, "%-20s\t%s\n", cmdPart, cmdShort); err != nil {
		return err
	}
	return nil
}

func printFlagLine(w io.Writer, flag *pflag.Flag) error {
	if flag.Hidden {
		return nil
	}

	var flagPart string
	if flag.Shorthand != "" {
		flagPart = fmt.Sprintf("  -%s, --%s", flag.Shorthand, flag.Name)
	} else {
		flagPart = fmt.Sprintf("  --%s", flag.Name)
	}

	if _, err := fmt.Fprintf(w, "%-20s\t%s\n", flagPart, flag.Usage); err != nil {
		return err
	}
	return nil
}
