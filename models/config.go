package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigStorage interface {
	Load(data interface{}) error
	Save(data interface{}) error
}

type FileConfig struct {
	Path string
}

func (f *FileConfig) Load(data interface{}) error {
	file, err := os.ReadFile(f.Path)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %v", f.Path, err)
	}

	if err := json.Unmarshal(file, data); err != nil {
		return fmt.Errorf("failed to parse config file %s: %v", f.Path, err)
	}

	return nil
}

func (f *FileConfig) Save(data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode config data: %v", err)
	}

	if err := os.MkdirAll("config", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	if err := os.WriteFile(f.Path, file, 0644); err != nil {
		return fmt.Errorf("failed to save config file %s: %v", f.Path, err)
	}

	return nil
}

var (
	ModelConfig   = &FileConfig{Path: "config/model.json"}
	PromptsConfig = &FileConfig{Path: "config/prompts.json"}
)
