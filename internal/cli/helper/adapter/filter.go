package adapter

import "github.com/spf13/cobra"

type FilterAdapter struct {
	command    *cobra.Command
	filterKeys []string
}

func NewFilterAdapter(cmd *cobra.Command) *FilterAdapter {
	return &FilterAdapter{command: cmd}
}

func (h *FilterAdapter) FilterKeys() []string {
	return h.filterKeys
}

func (h *FilterAdapter) AddFilterKey(key string) {
	if key != "" {
		h.filterKeys = append(h.filterKeys, key)
	}
}

func (h *FilterAdapter) AddFilterKeys(keys ...string) {
	for i := range keys {
		h.AddFilterKey(keys[i])
	}
}
