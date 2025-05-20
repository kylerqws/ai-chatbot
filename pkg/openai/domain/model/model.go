package model

import (
	"fmt"
	"strings"
)

type Model struct {
	Code        string
	Description string
}

const (
	GPT35Code = "gpt-3.5-turbo"
	GPT4Code  = "gpt-4"
)

var (
	GPT35 = &Model{
		Code:        GPT35Code,
		Description: "Fast and cost-effective model for general tasks and fine-tuning.",
	}
	GPT4 = &Model{
		Code:        GPT4Code,
		Description: "High-performance model for complex tasks, not fine-tunable.",
	}
)

var AllModels = map[string]*Model{
	GPT35Code: GPT35,
	GPT4Code:  GPT4,
}

func Resolve(code string) (*Model, error) {
	if code == "" {
		return GPT35, nil
	}
	if model, ok := AllModels[code]; ok {
		return model, nil
	}

	return nil, fmt.Errorf("unknown value '%v'", code)
}

func JoinCodes(sep string) string {
	codes := make([]string, 0, len(AllModels))
	for code := range AllModels {
		codes = append(codes, code)
	}

	return strings.Join(codes, sep)
}
