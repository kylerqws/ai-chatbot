package adapter

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	DateExample     = "1970-01-01"
	DatetimeExample = "1970-01-01 00:00:00"
)

type DateTimeAdapter struct {
	command         *cobra.Command
	datetimeFormats []string
}

func NewDateTimeAdapter(cmd *cobra.Command) *DateTimeAdapter {
	return &DateTimeAdapter{command: cmd, datetimeFormats: []string{time.DateOnly, time.DateTime}}
}

func (h *DateTimeAdapter) ParseDateTime(dateStr string) int64 {
	if strings.TrimSpace(dateStr) == "" {
		return 0
	}

	now := time.Now()
	loc := now.Location()

	for i := range h.datetimeFormats {
		t, err := time.ParseInLocation(h.datetimeFormats[i], dateStr, loc)
		if err == nil {
			return t.UTC().Unix()
		}
	}

	return now.UTC().Unix()
}

func (*DateTimeAdapter) NowDate() string {
	return time.Now().UTC().Format(time.DateOnly)
}

func (*DateTimeAdapter) Date(years, months, days int) string {
	return time.Now().AddDate(years, months, days).UTC().Format(time.DateOnly)
}

func (*DateTimeAdapter) NowDateTime() string {
	return time.Now().UTC().Format(time.DateTime)
}

func (*DateTimeAdapter) DateTime(years, months, days int) string {
	return time.Now().AddDate(years, months, days).UTC().Format(time.DateTime)
}

func (h *DateTimeAdapter) ValidateDateFormat(dateStr string) error {
	var err error
	for i := range h.datetimeFormats {
		if _, err = time.Parse(h.datetimeFormats[i], dateStr); err == nil {
			return nil
		}
	}
	return fmt.Errorf("invalid date format in '%s': %w", dateStr, err)
}
