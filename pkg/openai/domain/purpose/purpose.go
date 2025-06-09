package purpose

import (
	"fmt"
	"sort"
	"strings"
)

// Purpose defines a file usage category in OpenAI operations.
type Purpose struct {
	Code        string // Unique identifier for the purpose.
	Description string // Human-readable explanation of the purpose.
}

// Purpose code constants.
const (
	FineTuneCode   = "fine‑tune"
	AssistantsCode = "assistants"
	BatchCode      = "batch"
	UserDataCode   = "user_data"
	VisionCode     = "vision"
	EvalsCode      = "evals"
)

// Predefined Purpose instances.
var (
	FineTune   = &Purpose{Code: FineTuneCode, Description: "Used for fine‑tuning a model."}
	Assistants = &Purpose{Code: AssistantsCode, Description: "Used with the Assistants API."}
	Batch      = &Purpose{Code: BatchCode, Description: "Used with the Batch API."}
	UserData   = &Purpose{Code: UserDataCode, Description: "Used for user‑provided data."}
	Vision     = &Purpose{Code: VisionCode, Description: "Used with vision‑related APIs."}
	Evals      = &Purpose{Code: EvalsCode, Description: "Used with OpenAI Evals."}
)

// AllPurposes lists all known Purpose instances.
var AllPurposes = map[string]*Purpose{
	FineTuneCode:   FineTune,
	AssistantsCode: Assistants,
	BatchCode:      Batch,
	UserDataCode:   UserData,
	VisionCode:     Vision,
	EvalsCode:      Evals,
}

// Resolve looks up a Purpose by code, defaulting to FineTune.
func Resolve(code string) (*Purpose, error) {
	if code == "" {
		return FineTune, nil
	}
	if pr, ok := AllPurposes[code]; ok {
		return pr, nil
	}
	return nil, fmt.Errorf("unknown purpose code: '%s'", code)
}

// JoinCodes returns all purpose codes joined by separator.
func JoinCodes(sep string) string {
	c := make([]string, 0, len(AllPurposes))
	for code := range AllPurposes {
		c = append(c, code)
	}
	sort.Strings(c)
	return strings.Join(c, sep)
}
