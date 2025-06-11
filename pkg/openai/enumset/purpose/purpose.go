package purpose

import "github.com/kylerqws/chatbot/pkg/openai/utils/enumset"

// Purpose defines a file usage category in OpenAI operations.
type Purpose struct {
	Code        string // Unique identifier for the purpose.
	Description string // Human-readable explanation of the purpose.
}

// Purpose code constants.
const (
	FineTuneCode   = "fine-tune"
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
	return enumset.ResolveWithDefault(code, AllPurposes, FineTune, "purpose")
}

// JoinCodes returns all purpose codes joined by separator.
func JoinCodes(sep string) string {
	return enumset.JoinCodes(AllPurposes, sep)
}
