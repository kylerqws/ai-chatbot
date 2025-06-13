package helper

import (
	"fmt"
	"time"
)

// FormatBytes converts a byte count into a human-readable string.
func FormatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}

	div, exp := float64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	units := []string{"k", "M", "G", "T", "P", "E"}
	return fmt.Sprintf("%.2f %sB", float64(b)/div, units[exp])
}

// FormatTime converts a Unix timestamp to a formatted date-time string.
func FormatTime(t int64) string {
	return time.Unix(t, 0).Format(time.DateTime)
}
