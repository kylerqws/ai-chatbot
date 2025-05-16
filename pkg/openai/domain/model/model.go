package model

import (
	"fmt"
	"strings"
)

type Model struct {
	Code        string
	Description string
}

var (
	GPT35 = Model{
		Code:        "gpt-3.5-turbo",
		Description: "Fast and cost-effective model for general tasks and fine-tuning.",
	}
	GPT4 = Model{
		Code:        "gpt-4",
		Description: "High-performance model for complex tasks, not fine-tunable.",
	}
)

var AllModels = map[string]*Model{
	GPT35.Code: &GPT35,
	GPT4.Code:  &GPT4,
}

func Resolve(code string) (*Model, error) {
	if code == "" {
		return &GPT35, nil
	}
	if mdl, ok := AllModels[code]; ok {
		return mdl, nil
	}

	return nil, fmt.Errorf("[model.Resolve] unknown value '%v'", code)
}

func JoinCodes(sep string) string {
	var codes []string
	for _, mdl := range AllModels {
		codes = append(codes, mdl.Code)
	}

	return strings.Join(codes, sep)
}
