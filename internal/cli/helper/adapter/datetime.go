package adapter

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var DateFormats = []string{time.DateOnly, time.DateTime}

type DateTimeAdapter struct {
	command *cobra.Command
}

func NewDateTimeAdapter(cmd *cobra.Command) *DateTimeAdapter {
	return &DateTimeAdapter{command: cmd}
}

func (*DateTimeAdapter) ParseDateTime(dateStr string) int64 {
	now := time.Now()
	loc := now.Location()

	for i := range DateFormats {
		t, err := time.ParseInLocation(DateFormats[i], dateStr, loc)
		if err == nil {
			return t.UTC().Unix()
		}
	}

	return now.UTC().Unix()
}

func (*DateTimeAdapter) Date(years, months, days int) string {
	return time.Now().AddDate(years, months, days).UTC().Format(time.DateOnly)
}

func (*DateTimeAdapter) DateTime(years, months, days int) string {
	return time.Now().AddDate(years, months, days).UTC().Format(time.DateTime)
}

func (*DateTimeAdapter) ValidateDateFormat(dateStr string) error {
	var err error
	for i := range DateFormats {
		if _, err = time.Parse(DateFormats[i], dateStr); err == nil {
			return nil
		}
	}
	return fmt.Errorf("invalid date format in '%s': %w", dateStr, err)
}
