package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
)

func ExportToJSON(results []analyzer.LogResult, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("impossible de créer le fichier de sortie: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("erreur lors de l'encodage JSON: %w", err)
	}

	return nil
}