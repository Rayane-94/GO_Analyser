package analyzer

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Erreurs personnalisées
type FileNotFoundError struct {
	Path string
}

func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable: %s", e.Path)
}

type ParseError struct {
	Path    string
	Details string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("erreur d'analyse du fichier %s: %s", e.Path, e.Details)
}

type LogResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

func AnalyzeLog(id, path string) LogResult {
	result := LogResult{
		LogID:    id,
		FilePath: path,
	}

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fileErr := FileNotFoundError{Path: path}
			result.Status = "FAILED"
			result.Message = "Fichier introuvable"
			result.ErrorDetails = fileErr.Error()
			return result
		}
		
		result.Status = "FAILED" 
		result.Message = "Fichier inaccessible"
		result.ErrorDetails = err.Error()
		return result
	}

	// Simulation d'analyse avec délai aléatoire
	sleepDuration := time.Duration(rand.Intn(151)+50) * time.Millisecond
	time.Sleep(sleepDuration)

	result.Status = "OK"
	result.Message = "Analyse terminée avec succès"
	result.ErrorDetails = ""
	
	return result
}

func IsFileNotFoundError(err error) bool {
	var fileErr FileNotFoundError
	return errors.As(err, &fileErr)
}

func IsParseError(err error) bool {
	var parseErr ParseError
	return errors.As(err, &parseErr)
}