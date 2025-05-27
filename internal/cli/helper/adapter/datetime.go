package adapter

import (
	"github.com/spf13/cobra"
	"time"
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

func (*DateTimeAdapter) DateTime(years, months, days int) string {
	return time.Now().AddDate(years, months, days).UTC().Format(time.DateTime)
}
