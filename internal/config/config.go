package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// LogConfig représente la configuration d'un log à analyser
// Cette structure correspond exactement au format JSON attendu
type LogConfig struct {
	ID   string `json:"id"`   // L'identifiant unique du log
	Path string `json:"path"` // Le chemin vers le fichier de log  
	Type string `json:"type"` // Le type de log (nginx, custom, etc.)
}

// LoadConfig lit le fichier de configuration JSON et retourne une liste de LogConfig
func LoadConfig(configPath string) ([]LogConfig, error) {
	// Ouvrir le fichier de configuration
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier de config: %w", err)
	}
	// Fermer le fichier quand la fonction se termine (defer = "fais ça à la fin")
	defer file.Close()

	// Créer une variable pour stocker notre configuration
	var configs []LogConfig

	// Décoder le JSON dans notre variable
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configs); err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du JSON: %w", err)
	}

	return configs, nil
}