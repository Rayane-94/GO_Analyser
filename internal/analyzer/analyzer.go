package analyzer

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Définition des erreurs personnalisées
// Ces erreurs nous permettront de savoir exactement ce qui s'est mal passé

// FileNotFoundError représente une erreur quand le fichier n'existe pas
type FileNotFoundError struct {
	Path string
}

// Error implémente l'interface error pour FileNotFoundError
func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable: %s", e.Path)
}

// ParseError représente une erreur de parsing/analyse
type ParseError struct {
	Path    string
	Details string
}

// Error implémente l'interface error pour ParseError
func (e ParseError) Error() string {
	return fmt.Sprintf("erreur d'analyse du fichier %s: %s", e.Path, e.Details)
}

// LogResult représente le résultat de l'analyse d'un log
type LogResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

// AnalyzeLog analyse un seul fichier de log
// Cette fonction simule une analyse avec une attente aléatoire
func AnalyzeLog(id, path string) LogResult {
	// Créer le résultat par défaut
	result := LogResult{
		LogID:    id,
		FilePath: path,
	}

	// Vérifier si le fichier existe et est lisible
	_, err := os.Stat(path)
	if err != nil {
		// Si le fichier n'existe pas, créer une erreur personnalisée
		if os.IsNotExist(err) {
			fileErr := FileNotFoundError{Path: path}
			result.Status = "FAILED"
			result.Message = "Fichier introuvable"
			result.ErrorDetails = fileErr.Error()
			return result
		}
		
		// Autre type d'erreur (permissions, etc.)
		result.Status = "FAILED" 
		result.Message = "Fichier inaccessible"
		result.ErrorDetails = err.Error()
		return result
	}

	// Simuler le temps d'analyse avec une durée aléatoire entre 50 et 200ms
	// rand.Intn(151) donne un nombre entre 0 et 150, on ajoute 50 pour avoir 50-200
	sleepDuration := time.Duration(rand.Intn(151)+50) * time.Millisecond
	time.Sleep(sleepDuration)

	// Si tout va bien, l'analyse est réussie
	result.Status = "OK"
	result.Message = "Analyse terminée avec succès"
	result.ErrorDetails = ""
	
	return result
}

// IsFileNotFoundError vérifie si une erreur est de type FileNotFoundError
// Utilise errors.As pour vérifier le type d'erreur
func IsFileNotFoundError(err error) bool {
	var fileErr FileNotFoundError
	return errors.As(err, &fileErr)
}

// IsParseError vérifie si une erreur est de type ParseError
func IsParseError(err error) bool {
	var parseErr ParseError
	return errors.As(err, &parseErr)
}