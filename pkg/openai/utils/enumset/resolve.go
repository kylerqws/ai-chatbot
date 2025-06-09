package enumset

import "fmt"

// ResolveWithDefault returns the item from the map or a default if code is empty.
func ResolveWithDefault[V any](code string, m map[string]V, defaultVal V, name string) (V, error) {
	if code == "" {
		return defaultVal, nil
	}
	if v, ok := m[code]; ok {
		return v, nil
	}
	var zero V
	return zero, fmt.Errorf("unknown %s code: '%s'", name, code)
}

// ResolveRequired returns the item from the map or an error if code is empty.
func ResolveRequired[V any](code string, m map[string]V, name string) (V, error) {
	if code == "" {
		var zero V
		return zero, fmt.Errorf("%s code is required", name)
	}
	if v, ok := m[code]; ok {
		return v, nil
	}
	var zero V
	return zero, fmt.Errorf("unknown %s code: '%s'", name, code)
}
