package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
)

// ExportToJSON exporte les résultats d'analyse vers un fichier JSON
func ExportToJSON(results []analyzer.LogResult, outputPath string) error {
	// Créer le fichier de sortie
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("impossible de créer le fichier de sortie: %w", err)
	}
	// S'assurer que le fichier sera fermé à la fin de la fonction
	defer file.Close()

	// Créer un encoder JSON avec une indentation pour rendre le fichier lisible
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Indentation avec 2 espaces

	// Encoder les résultats en JSON dans le fichier
	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("erreur lors de l'encodage JSON: %w", err)
	}

	return nil
}