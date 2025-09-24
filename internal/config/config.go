package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func LoadConfig(configPath string) ([]LogConfig, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier de config: %w", err)
	}
	defer file.Close()

	var configs []LogConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configs); err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du JSON: %w", err)
	}

	return configs, nil
}