package adapter

import ctr "github.com/kylerqws/chatbot/internal/cli/contract"

// ValidateAdapter defines the interface for validating CLI flags and arguments.
type ValidateAdapter interface {
	// ValidateRequireFlag checks if a required flag is provided.
	ValidateRequireFlag(name string) error

	// ValidateHasChangedAnyFlag checks if any of the specified flags were provided.
	ValidateHasChangedAnyFlag(names ...string) error

	// ValidateHasMoreArgsThan checks if argument count exceeds the limit.
	ValidateHasMoreArgsThan(count uint8) error

	// ValidateStringFlag validates a string flag.
	ValidateStringFlag(name string, fn ctr.FuncValidateString) error

	// ValidateStringSliceFlag validates a string slice flag.
	ValidateStringSliceFlag(name string, fn ctr.FuncValidateString) []error

	// ValidateUint8Flag validates a uint8 flag.
	ValidateUint8Flag(name string, fn ctr.FuncValidateUint8) error

	// ValidateUint8SliceFlag validates a uint8 slice flag.
	ValidateUint8SliceFlag(name string, fn ctr.FuncValidateUint8) []error

	// ValidateUintFlag validates a uint flag.
	ValidateUintFlag(name string, fn ctr.FuncValidateUint) error

	// ValidateUintSliceFlag validates a uint slice flag.
	ValidateUintSliceFlag(name string, fn ctr.FuncValidateUint) []error
}
