package adapter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	AfterFlagKey = "after"
	LimitFlagKey = "limit"
)

const (
	DefaultLimit = 20
	MinimumLimit = 1
	MaximumLimit = 100
)

type PaginationAdapter struct {
	command *cobra.Command
}

func NewPaginationAdapter(cmd *cobra.Command) *PaginationAdapter {
	return &PaginationAdapter{command: cmd}
}

func (*FormatAdapter) LimitToString(limit uint8) string {
	return strconv.Itoa(int(limit))
}

func (h *FormatAdapter) JoinLimits(sep string) string {
	return strings.Join([]string{
		"default " + h.LimitToString(DefaultLimit),
		"minimum " + h.LimitToString(MinimumLimit),
		"maximum " + h.LimitToString(MaximumLimit),
	}, sep)
}

func (*PaginationAdapter) ValidateLimit(limit uint8) error {
	if limit < MinimumLimit {
		return fmt.Errorf("limit must be ≥ %d", MinimumLimit)
	}
	if limit > MaximumLimit {
		return fmt.Errorf("limit must be ≤ %d", MaximumLimit)
	}

	return nil
}
