package setup

import "github.com/spf13/cobra"

func AddStringFlag(cmd *cobra.Command, name, shorthand, value, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().StringP(name, shorthand, value, desc)
	} else {
		cmd.Flags().StringP(name, shorthand, value, desc)
	}
}

func AddStringSliceFlag(cmd *cobra.Command, name, shorthand string, value []string, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().StringSliceP(name, shorthand, value, desc)
	} else {
		cmd.Flags().StringSliceP(name, shorthand, value, desc)
	}
}

func AddBoolFlag(cmd *cobra.Command, name, shorthand string, value bool, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().BoolP(name, shorthand, value, desc)
	} else {
		cmd.Flags().BoolP(name, shorthand, value, desc)
	}
}
