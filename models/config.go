package models

import (
	"encoding/json"
	"fmt"
	"os"
)

const modelConfigPath = "config/model.json"

type ModelData struct {
	Name   string `json:"name"`
	FileID string `json:"file_id"`
}

var ModelConfig = &ConfigManager{Path: modelConfigPath}

type ConfigManager struct {
	Path string
}

func (c *ConfigManager) Load() (*ModelData, error) {
	file, err := os.Open(c.Path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open config file: %v", err)
	}
	defer file.Close()

	var modelData ModelData
	if err := json.NewDecoder(file).Decode(&modelData); err != nil {
		return nil, fmt.Errorf("Failed to parse config file: %v", err)
	}

	return &modelData, nil
}

func (c *ConfigManager) Save(modelData *ModelData) error {
	file, err := os.Create(c.Path)
	if err != nil {
		return fmt.Errorf("Failed to create config file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(modelData); err != nil {
		return fmt.Errorf("Failed to write config file: %v", err)
	}

	return nil
}
