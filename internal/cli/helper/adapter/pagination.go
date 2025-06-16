package adapter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	// AfterFlagKey is the flag name for pagination "after".
	AfterFlagKey = "after"

	// LimitFlagKey is the flag name for pagination "limit".
	LimitFlagKey = "limit"
)

const (
	// DefaultLimit is the default pagination limit.
	DefaultLimit = 20

	// MinimumLimit is the minimum allowed limit.
	MinimumLimit = 1

	// MaximumLimit is the maximum allowed limit.
	MaximumLimit = 100
)

// PaginationAdapter provides the implementation for CLI adapter with pagination handling.
type PaginationAdapter struct {
	command *cobra.Command
}

// NewPaginationAdapter creates a new instance of PaginationAdapter.
func NewPaginationAdapter(cmd *cobra.Command) *PaginationAdapter {
	return &PaginationAdapter{command: cmd}
}

// LimitToString converts a limit value to a string.
func (*PaginationAdapter) LimitToString(limit uint8) string {
	return strconv.Itoa(int(limit))
}

// JoinLimits returns the default, min and max limits as a formatted string.
func (a *PaginationAdapter) JoinLimits(sep string) string {
	return strings.Join([]string{
		a.LimitToString(MinimumLimit) + "–" + a.LimitToString(MaximumLimit),
		"default " + a.LimitToString(DefaultLimit),
	}, sep)
}

// ValidateLimit validates that the limit is within allowed bounds.
func (*PaginationAdapter) ValidateLimit(limit uint8) error {
	if limit < MinimumLimit {
		return fmt.Errorf("limit must be ≥ %d", MinimumLimit)
	}
	if limit > MaximumLimit {
		return fmt.Errorf("limit must be ≤ %d", MaximumLimit)
	}
	return nil
}
