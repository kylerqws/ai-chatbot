package value

import (
	"encoding/json"
	"fmt"
)

// AutoOrNumber is a generic type representing either the string "auto"
// or a numeric value of type T (int or float64). Commonly used in API parameters
// that support automatic or explicit numeric configuration.
//
// Examples:
//
//	"auto"
//	42
//	3.14
type AutoOrNumber[T int | float64] struct {
	// IsAuto indicates whether the value is set to "auto".
	IsAuto bool

	// Value holds the numeric value if IsAuto is false.
	Value T
}

// NewAuto returns a new AutoOrNumber set to "auto".
func NewAuto[T int | float64]() *AutoOrNumber[T] {
	return &AutoOrNumber[T]{IsAuto: true}
}

// NewNumber returns a new AutoOrNumber set to the provided numeric value.
func NewNumber[T int | float64](value T) *AutoOrNumber[T] {
	return &AutoOrNumber[T]{IsAuto: false, Value: value}
}

// MarshalJSON serializes the value as either the string "auto" or the numeric value.
func (v *AutoOrNumber[T]) MarshalJSON() ([]byte, error) {
	if v.IsAuto {
		return json.Marshal("auto")
	}
	return json.Marshal(v.Value)
}

// UnmarshalJSON deserializes the value from either the string "auto" or a numeric value.
func (v *AutoOrNumber[T]) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if s == "auto" {
			v.IsAuto = true
			var zero T
			v.Value = zero
			return nil
		}
		return fmt.Errorf("invalid string value for AutoOrNumber: '%s'", s)
	}

	var num T
	if err := json.Unmarshal(data, &num); err != nil {
		return fmt.Errorf("parse numeric value: %w", err)
	}

	v.Value = num
	v.IsAuto = false
	return nil
}

// IsSetAuto reports whether the value is "auto".
func (v *AutoOrNumber[T]) IsSetAuto() bool {
	return v.IsAuto
}

// GetValue returns the numeric value and a boolean indicating if it is explicitly set.
func (v *AutoOrNumber[T]) GetValue() (T, bool) {
	return v.Value, !v.IsAuto
}
