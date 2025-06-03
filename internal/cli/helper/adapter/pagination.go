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

func (*PaginationAdapter) MaximumLimit() uint8 {
	return MaximumLimit
}

func (*PaginationAdapter) StrMaximumLimit() string {
	return strconv.FormatUint(uint64(MaximumLimit), 10)
}

func (*PaginationAdapter) MinimumLimit() uint8 {
	return MinimumLimit
}

func (*PaginationAdapter) StrMinimumLimit() string {
	return strconv.FormatUint(uint64(MinimumLimit), 10)
}

func (*PaginationAdapter) DefaultLimit() uint8 {
	return DefaultLimit
}

func (*PaginationAdapter) StrDefaultLimit() string {
	return strconv.FormatUint(uint64(DefaultLimit), 10)
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
