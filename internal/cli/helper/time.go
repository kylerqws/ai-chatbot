package helper

import (
	"time"

	"github.com/spf13/cobra"
)

type DateTimeAdapterHelper struct {
	command *cobra.Command
}

func NewDateTimeAdapterHelper(cmd *cobra.Command) *DateTimeAdapterHelper {
	return &DateTimeAdapterHelper{command: cmd}
}

func (*DateTimeAdapterHelper) ParseDateTime(dateStr string) int64 {
	now := time.Now()
	loc := now.Location()

	formats := []string{time.DateOnly, time.DateTime}
	for _, layout := range formats {
		if t, err := time.ParseInLocation(layout, dateStr, loc); err == nil {
			return t.UTC().Unix()
		}
	}

	return now.UTC().Unix()
}

func (*DateTimeAdapterHelper) DateTime(years, months, days int) string {
	return time.Now().AddDate(years, months, days).UTC().Format(time.DateTime)
}
