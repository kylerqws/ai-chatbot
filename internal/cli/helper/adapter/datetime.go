package adapter

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	// DateExample is an example of a valid date string.
	DateExample = "2025-03-17"

	// DatetimeExample is an example of a valid datetime string.
	DatetimeExample = "2025-03-17 14:45:00"
)

// DateTimeAdapter provides the implementation of a CLI adapter for date and time handling.
type DateTimeAdapter struct {
	command         *cobra.Command
	datetimeFormats []string
}

// NewDateTimeAdapter creates a new instance of DateTimeAdapter.
func NewDateTimeAdapter(cmd *cobra.Command) *DateTimeAdapter {
	return &DateTimeAdapter{command: cmd, datetimeFormats: []string{time.DateOnly, time.DateTime}}
}

// ParseDateTime parses a date or datetime string into a UTC Unix timestamp.
func (a *DateTimeAdapter) ParseDateTime(dateStr string) *int64 {
	if strings.TrimSpace(dateStr) == "" {
		return a.datetimePointer(0)
	}

	for i := range a.datetimeFormats {
		t, err := time.ParseInLocation(a.datetimeFormats[i], dateStr, time.UTC)
		if err == nil {
			return a.datetimePointer(t.UTC().Unix())
		}
	}

	return a.datetimePointer(time.Now().UTC().Unix())
}

// NowDate returns the current UTC date.
func (*DateTimeAdapter) NowDate() string {
	return time.Now().UTC().Format(time.DateOnly)
}

// Date returns a UTC date offset by the given values.
func (*DateTimeAdapter) Date(years, months, days int) string {
	return time.Now().AddDate(years, months, days).UTC().Format(time.DateOnly)
}

// NowDateTime returns the current UTC datetime.
func (*DateTimeAdapter) NowDateTime() string {
	return time.Now().UTC().Format(time.DateTime)
}

// DateTime returns a UTC datetime offset by the given values.
func (*DateTimeAdapter) DateTime(years, months, days int) string {
	return time.Now().AddDate(years, months, days).UTC().Format(time.DateTime)
}

// ValidateDateFormat validates a date or datetime string.
func (a *DateTimeAdapter) ValidateDateFormat(dateStr string) error {
	var err error
	for i := range a.datetimeFormats {
		if _, err = time.Parse(a.datetimeFormats[i], dateStr); err == nil {
			return nil
		}
	}
	return fmt.Errorf("invalid date or datetime format in '%s': %w", dateStr, err)
}

// datetimePointer returns a pointer to the value or nil if it's zero.
func (*DateTimeAdapter) datetimePointer(val int64) *int64 {
	if val == 0 {
		return nil
	}
	return &val
}
