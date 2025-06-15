package setup

import "github.com/spf13/cobra"

// AddStringFlag adds a string flag to the given command.
// The inherited argument specifies whether the flag is persistent (true) or local (false).
func AddStringFlag(cmd *cobra.Command, name, shorthand, value, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().StringP(name, shorthand, value, desc)
	} else {
		cmd.Flags().StringP(name, shorthand, value, desc)
	}
}

// AddStringSliceFlag adds a string slice flag to the given command.
// The inherited argument specifies whether the flag is persistent (true) or local (false).
func AddStringSliceFlag(cmd *cobra.Command, name, shorthand string, value []string, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().StringSliceP(name, shorthand, value, desc)
	} else {
		cmd.Flags().StringSliceP(name, shorthand, value, desc)
	}
}

// AddUint8Flag adds an uint8 flag to the given command.
// The inherited argument specifies whether the flag is persistent (true) or local (false).
func AddUint8Flag(cmd *cobra.Command, name, shorthand string, value uint8, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().Uint8P(name, shorthand, value, desc)
	} else {
		cmd.Flags().Uint8P(name, shorthand, value, desc)
	}
}

// AddUint8SliceFlag adds an uint8 slice flag to the given command.
// The inherited argument specifies whether the flag is persistent (true) or local (false).
func AddUint8SliceFlag(_ *cobra.Command, _, _ string, _ []uint8, _ string, _ bool) {
	panic("uint8 slice flag is not implemented")
}

// AddUintFlag adds an uint flag to the given command.
// The inherited argument specifies whether the flag is persistent (true) or local (false).
func AddUintFlag(cmd *cobra.Command, name, shorthand string, value uint, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().UintP(name, shorthand, value, desc)
	} else {
		cmd.Flags().UintP(name, shorthand, value, desc)
	}
}

// AddUintSliceFlag adds an uint slice flag to the given command.
// The inherited argument specifies whether the flag is persistent (true) or local (false).
func AddUintSliceFlag(cmd *cobra.Command, name, shorthand string, value []uint, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().UintSliceP(name, shorthand, value, desc)
	} else {
		cmd.Flags().UintSliceP(name, shorthand, value, desc)
	}
}

// AddBoolFlag adds a boolean flag to the given command.
// The inherited argument specifies whether the flag is persistent (true) or local (false).
func AddBoolFlag(cmd *cobra.Command, name, shorthand string, value bool, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().BoolP(name, shorthand, value, desc)
	} else {
		cmd.Flags().BoolP(name, shorthand, value, desc)
	}
}
