package adapter

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

const (
	MinimumLimit = 1
	MaximumLimit = 100
	DefaultLimit = 20
)

type limitValue struct {
	val uint8
	str string
}

type limit struct {
	max *limitValue
	min *limitValue
	def *limitValue
}

type PaginationAdapter struct {
	command *cobra.Command
	limit   *limit
}

func NewPaginationAdapter(cmd *cobra.Command) *PaginationAdapter {
	return &PaginationAdapter{
		command: cmd,
		limit:   &limit{max: &limitValue{}, min: &limitValue{}, def: &limitValue{}},
	}
}

func (a *PaginationAdapter) MaximumLimit() uint8 {
	if a.limit.max.val == 0 {
		a.limit.max.val = MaximumLimit
	}
	return a.limit.max.val
}

func (a *PaginationAdapter) StrMaximumLimit() string {
	if a.limit.max.str == "" {
		a.limit.max.str = strconv.FormatUint(uint64(a.MaximumLimit()), 10)
	}
	return a.limit.max.str
}

func (a *PaginationAdapter) MinimumLimit() uint8 {
	if a.limit.min.val == 0 {
		a.limit.min.val = MinimumLimit
	}
	return a.limit.min.val
}

func (a *PaginationAdapter) StrMinimumLimit() string {
	if a.limit.min.str == "" {
		a.limit.min.str = strconv.FormatUint(uint64(a.MinimumLimit()), 10)
	}
	return a.limit.min.str
}

func (a *PaginationAdapter) DefaultLimit() uint8 {
	if a.limit.def.val == 0 {
		a.limit.def.val = DefaultLimit
	}
	return a.limit.def.val
}

func (a *PaginationAdapter) StrDefaultLimit() string {
	if a.limit.def.str == "" {
		a.limit.def.str = strconv.FormatUint(uint64(a.DefaultLimit()), 10)
	}
	return a.limit.def.str
}

func (a *PaginationAdapter) ValidateLimit(limit uint8) error {
	maxLimit, minLimit := a.MaximumLimit(), a.MinimumLimit()

	if limit > maxLimit {
		return fmt.Errorf("limit must be ≤ %d", maxLimit)
	}
	if limit < minLimit {
		return fmt.Errorf("limit must be ≥ %d", minLimit)
	}

	return nil
}
