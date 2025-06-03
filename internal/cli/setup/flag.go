package setup

import (
	"github.com/spf13/cobra"
	"log"
)

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

func AddUint8Flag(cmd *cobra.Command, name, shorthand string, value uint8, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().Uint8P(name, shorthand, value, desc)
	} else {
		cmd.Flags().Uint8P(name, shorthand, value, desc)
	}
}

func AddUint8SliceFlag(_ *cobra.Command, _, _ string, _ []uint8, _ string, _ bool) {
	// TODO: need to implemented Uint8P and Uint8SliceP as they are not implemented in the Cobra package
	log.Fatalf("method 'AddUint8SliceFlag' is not implemented")
}

func AddUintFlag(cmd *cobra.Command, name, shorthand string, value uint, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().UintP(name, shorthand, value, desc)
	} else {
		cmd.Flags().UintP(name, shorthand, value, desc)
	}
}

func AddUintSliceFlag(cmd *cobra.Command, name, shorthand string, value []uint, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().UintSliceP(name, shorthand, value, desc)
	} else {
		cmd.Flags().UintSliceP(name, shorthand, value, desc)
	}
}

func AddBoolFlag(cmd *cobra.Command, name, shorthand string, value bool, desc string, inherited bool) {
	if inherited {
		cmd.PersistentFlags().BoolP(name, shorthand, value, desc)
	} else {
		cmd.Flags().BoolP(name, shorthand, value, desc)
	}
}
