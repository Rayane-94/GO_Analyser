package cmd

import (
	"fmt"
	"sync"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
	"github.com/axellelanca/go_loganizer/internal/config"
	"github.com/axellelanca/go_loganizer/internal/reporter"
	"github.com/spf13/cobra"
)

var (
	configPath string
	outputPath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les fichiers de logs spécifiés dans la configuration",
	Long: `La commande analyze lit un fichier de configuration JSON contenant 
les logs à analyser, puis traite chaque log en parallèle en utilisant 
des goroutines.

Exemple:
  loganalyzer analyze -c config.json -o report.json`,
	Run: runAnalyze,
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Chemin vers le fichier de configuration JSON (obligatoire)")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Chemin vers le fichier de sortie JSON (optionnel)")
	analyzeCmd.MarkFlagRequired("config")
}

func runAnalyze(cmd *cobra.Command, args []string) {
	fmt.Printf("Chargement de la configuration depuis: %s\n", configPath)
	
	configs, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Erreur lors du chargement de la configuration: %v\n", err)
		return
	}

	fmt.Printf("Configuration chargée: %d logs à analyser\n", len(configs))

	var wg sync.WaitGroup
	resultsChan := make(chan analyzer.LogResult, len(configs))

	for _, logConfig := range configs {
		wg.Add(1)
		go func(lc config.LogConfig) {
			defer wg.Done()
			fmt.Printf("Début de l'analyse: %s (%s)\n", lc.ID, lc.Path)
			result := analyzer.AnalyzeLog(lc.ID, lc.Path)
			resultsChan <- result
			
			if result.Status == "OK" {
				fmt.Printf("✓ %s: %s\n", result.LogID, result.Message)
			} else {
				fmt.Printf("✗ %s: %s - %s\n", result.LogID, result.Message, result.ErrorDetails)
			}
		}(logConfig)
	}

	wg.Wait()
	close(resultsChan)

	var results []analyzer.LogResult
	for result := range resultsChan {
		results = append(results, result)
	}

	fmt.Printf("\n=== RÉSUMÉ ===\n")
	fmt.Printf("Total des logs analysés: %d\n", len(results))
	
	successCount := 0
	failCount := 0
	for _, result := range results {
		if result.Status == "OK" {
			successCount++
		} else {
			failCount++
		}
	}
	
	fmt.Printf("Succès: %d\n", successCount)
	fmt.Printf("Échecs: %d\n", failCount)

	if outputPath != "" {
		fmt.Printf("\nExport des résultats vers: %s\n", outputPath)
		
		err := reporter.ExportToJSON(results, outputPath)
		if err != nil {
			fmt.Printf("Erreur lors de l'export: %v\n", err)
			return
		}
		
		fmt.Println("Export terminé avec succès!")
	}
}