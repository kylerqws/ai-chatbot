package purpose

import (
	"fmt"
	"strings"
)

type Purpose struct {
	Code        string
	Description string
}

const (
	FineTuneCode   = "fine-tune"
	AssistantsCode = "assistants"
	BatchCode      = "batch"
	UserDataCode   = "user_data"
	VisionCode     = "vision"
	EvalsCode      = "evals"
)

var (
	FineTune = &Purpose{
		Code:        FineTuneCode,
		Description: "Used to upload files for fine-tuning models.",
	}
	Assistants = &Purpose{
		Code:        AssistantsCode,
		Description: "Used to upload files for the Assistants API.",
	}
	Batch = &Purpose{
		Code:        BatchCode,
		Description: "Used for batch processing tasks.",
	}
	UserData = &Purpose{
		Code:        UserDataCode,
		Description: "Used to upload user data files.",
	}
	Vision = &Purpose{
		Code:        VisionCode,
		Description: "Used for vision model tasks.",
	}
	Evals = &Purpose{
		Code:        EvalsCode,
		Description: "Used for evaluation tasks.",
	}
)

var AllPurposes = map[string]*Purpose{
	FineTuneCode:   FineTune,
	AssistantsCode: Assistants,
	BatchCode:      Batch,
	UserDataCode:   UserData,
	VisionCode:     Vision,
	EvalsCode:      Evals,
}

func Resolve(code string) (*Purpose, error) {
	if code == "" {
		return FineTune, nil
	}
	if prp, ok := AllPurposes[code]; ok {
		return prp, nil
	}
	return nil, fmt.Errorf("unknown value '%v'", code)
}

func JoinCodes(sep string) string {
	codes := make([]string, 0, len(AllPurposes))
	for code := range AllPurposes {
		codes = append(codes, code)
	}
	return strings.Join(codes, sep)
}
