package eventlevel

import "github.com/kylerqws/chatbot/pkg/openai/utils/enumset"

// EventLevel defines the severity level of a fine‑tuning event.
type EventLevel struct {
	Code        string // Unique identifier for the event level.
	Description string // Human-readable explanation of the level.
}

// EventLevel code constants.
const (
	InfoCode    = "info"
	WarningCode = "warning"
	ErrorCode   = "error"
)

// Predefined EventLevel instances.
var (
	Info    = &EventLevel{Code: InfoCode, Description: "Informational event."}
	Warning = &EventLevel{Code: WarningCode, Description: "Potential issue."}
	Error   = &EventLevel{Code: ErrorCode, Description: "Error occurred."}
)

// AllEventLevels lists all known EventLevel instances.
var AllEventLevels = map[string]*EventLevel{
	InfoCode:    Info,
	WarningCode: Warning,
	ErrorCode:   Error,
}

// Resolve looks up an EventLevel by code, error if missing or unknown.
func Resolve(code string) (*EventLevel, error) {
	return enumset.ResolveRequired(code, AllEventLevels, "event level")
}

// JoinCodes returns all event level codes joined by separator.
func JoinCodes(sep string) string {
	return enumset.JoinCodes(AllEventLevels, sep)
}
