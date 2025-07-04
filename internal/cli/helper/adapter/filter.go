package adapter

import "github.com/spf13/cobra"

// FilterAdapter provides the implementation for CLI adapter with filter handling.
type FilterAdapter struct {
	command    *cobra.Command
	filterKeys []string
}

// NewFilterAdapter creates a new instance of FilterAdapter.
func NewFilterAdapter(cmd *cobra.Command) *FilterAdapter {
	return &FilterAdapter{command: cmd}
}

// FilterKeys returns the list of collected filter keys.
func (a *FilterAdapter) FilterKeys() []string {
	return a.filterKeys
}

// ExistFilterKeys reports whether any filter keys have been added.
func (a *FilterAdapter) ExistFilterKeys() bool {
	return len(a.filterKeys) > 0
}

// AddFilterKey adds a single filter key to the collection.
func (a *FilterAdapter) AddFilterKey(key string) {
	if key != "" {
		a.filterKeys = append(a.filterKeys, key)
	}
}

// AddFilterKeys adds multiple filter keys to the collection.
func (a *FilterAdapter) AddFilterKeys(keys ...string) {
	for i := range keys {
		a.AddFilterKey(keys[i])
	}
}
