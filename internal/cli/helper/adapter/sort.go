package adapter

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const (
	// SortOrderFlagKey is the flag name used for sorting.
	SortOrderFlagKey = "sort"
)

const (
	// SortOrderAsc and SortOrderDesc define supported sort order values.
	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"

	// DefaultSortOrder is the default sort order value.
	DefaultSortOrder = SortOrderAsc
)

const (
	// SortOrderAscExample is an example for ascending order.
	SortOrderAscExample = "ASC"

	// SortOrderDescExample is an example for descending order.
	SortOrderDescExample = "DESC"
)

// SortAdapter provides the implementation for CLI adapter with sort order handling.
type SortAdapter struct {
	command *cobra.Command
}

// NewSortAdapter creates a new instance of SortAdapter.
func NewSortAdapter(cmd *cobra.Command) *SortAdapter {
	return &SortAdapter{command: cmd}
}

// SortOrderToLower converts the sort order to lowercase.
func (*SortAdapter) SortOrderToLower(order *string) *string {
	res := strings.TrimSpace(strings.ToLower(*order))
	return &res
}

// NormalizeSortOrder returns the normalized sort order value or the default if empty.
func (a *SortAdapter) NormalizeSortOrder(order string) string {
	orderLower := a.SortOrderToLower(&order)
	if *orderLower == "" {
		return DefaultSortOrder
	}
	return *orderLower
}

// ValidateSortOrder checks whether the given sort order is valid.
func (a *SortAdapter) ValidateSortOrder(order string) error {
	order = a.NormalizeSortOrder(order)
	if order == SortOrderAsc || order == SortOrderDesc {
		return nil
	}
	return fmt.Errorf("invalid sort order '%s'", order)
}
