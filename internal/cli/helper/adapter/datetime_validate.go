package adapter

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type ValidateDateTimeAdapter struct {
	*DateTimeAdapter
	*ValidateFlagAdapter
	command *cobra.Command
}

func NewValidateDateTimeAdapter(cmd *cobra.Command) *ValidateDateTimeAdapter {
	hlp := &ValidateDateTimeAdapter{command: cmd}

	hlp.DateTimeAdapter = NewDateTimeAdapter(cmd)
	hlp.ValidateFlagAdapter = NewValidateFlagAdapter(cmd)

	return hlp
}

func (h *ValidateDateTimeAdapter) ValidateDateFlag(flagKey string) error {
	return h.ValidateStringFlag(flagKey, h.ValidateDateFormat)
}

func (*DateTimeAdapter) ValidateDateFormat(dateStr string) error {
	var err error
	for i := range DateFormats {
		if _, err = time.Parse(DateFormats[i], dateStr); err == nil {
			return nil
		}
	}
	return fmt.Errorf("invalid date format in %q: %w", dateStr, err)
}
