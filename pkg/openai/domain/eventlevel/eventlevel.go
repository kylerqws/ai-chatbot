package eventlevel

import (
	"fmt"
	"sort"
	"strings"
)

// EventLevel defines the severity level of a fine-tuning event.
type EventLevel struct {
	Code        string // Unique identifier for the event level.
	Description string // Human-readable explanation of the event level.
}

// Event level code constants.
const (
	InfoCode    = "info"
	WarningCode = "warning"
	ErrorCode   = "error"
)

// Predefined EventLevel instances.
var (
	Info = &EventLevel{
		Code:        InfoCode,
		Description: "Informational event.",
	}
	Warning = &EventLevel{
		Code:        WarningCode,
		Description: "Potential issue or caution.",
	}
	Error = &EventLevel{
		Code:        ErrorCode,
		Description: "Error that occurred during the process.",
	}
)

// AllEventLevels maps all known event level codes to their EventLevel instances.
var AllEventLevels = map[string]*EventLevel{
	InfoCode:    Info,
	WarningCode: Warning,
	ErrorCode:   Error,
}

// Resolve returns the EventLevel associated with the given code.
// Returns an error if the code is empty or unrecognized.
func Resolve(code string) (*EventLevel, error) {
	if code == "" {
		return nil, fmt.Errorf("event level code is required")
	}
	if evl, ok := AllEventLevels[code]; ok {
		return evl, nil
	}
	return nil, fmt.Errorf("unknown event level code: '%v'", code)
}

// JoinCodes returns a sorted, delimited string of all known event level codes.
func JoinCodes(sep string) string {
	codes := make([]string, 0, len(AllEventLevels))
	for code := range AllEventLevels {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return strings.Join(codes, sep)
}
