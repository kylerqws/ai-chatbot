package purpose

import (
	"fmt"
	"strings"
)

type Purpose struct {
	Code        string
	Description string
}

var (
	FineTune = Purpose{
		Code:        "fine-tune",
		Description: "Used to upload files for fine-tuning models.",
	}
	Assistants = Purpose{
		Code:        "assistants",
		Description: "Used to upload files for the Assistants API.",
	}
	Batch = Purpose{
		Code:        "batch",
		Description: "Used for batch processing tasks.",
	}
	UserData = Purpose{
		Code:        "user_data",
		Description: "Used to upload user data files.",
	}
	Vision = Purpose{
		Code:        "vision",
		Description: "Used for vision model tasks.",
	}
	Evals = Purpose{
		Code:        "evals",
		Description: "Used for evaluation tasks.",
	}
)

var AllPurposes = map[string]*Purpose{
	FineTune.Code:   &FineTune,
	Assistants.Code: &Assistants,
	Batch.Code:      &Batch,
	UserData.Code:   &UserData,
	Vision.Code:     &Vision,
	Evals.Code:      &Evals,
}

func Resolve(code string) (*Purpose, error) {
	if code == "" {
		return &FineTune, nil
	}
	if prp, ok := AllPurposes[code]; ok {
		return prp, nil
	}

	return nil, fmt.Errorf("[purpose.Resolve] unknown value '%v'", code)
}

func JoinCodes(sep string) string {
	var codes []string
	for _, prp := range AllPurposes {
		codes = append(codes, prp.Code)
	}

	return strings.Join(codes, sep)
}
