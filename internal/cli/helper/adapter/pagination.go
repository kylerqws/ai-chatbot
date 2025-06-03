package adapter

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

const (
	AfterFlagKey = "after"
	LimitFlagKey = "limit"
)

const (
	MinimumLimit = 1
	MaximumLimit = 100
	DefaultLimit = 20
)

type PaginationAdapter struct {
	command *cobra.Command
}

func NewPaginationAdapter(cmd *cobra.Command) *PaginationAdapter {
	return &PaginationAdapter{command: cmd}
}

func (*FormatAdapter) LimitToString(num uint8) string {
	return strconv.Itoa(int(num))
}

func (*PaginationAdapter) ValidateLimit(limit uint8) error {
	if limit > MaximumLimit {
		return fmt.Errorf("limit must be ≤ %d", MaximumLimit)
	}
	if limit < MinimumLimit {
		return fmt.Errorf("limit must be ≥ %d", MinimumLimit)
	}

	return nil
}
