package helper

import (
	"fmt"
	"strings"
	"time"
)

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

func FormatTime(t int64) string {
	return time.Unix(t, 0).Format(time.DateTime)
}

func FormatPadToWidth(s string, width int) string {
	if len(s) >= width {
		return s
	}

	totalPad := width - len(s)
	if totalPad%2 != 0 {
		totalPad--
	}

	pad := totalPad / 2
	return strings.Repeat(" ", pad) + s + strings.Repeat(" ", pad)
}

func FormatPadCenter(s string, padding int) string {
	if padding <= 0 {
		return s
	}
	pad := strings.Repeat(" ", padding)

	return pad + s + pad
}
