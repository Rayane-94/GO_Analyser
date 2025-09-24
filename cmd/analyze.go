package cmd

import (
	"fmt"
	"sync"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
	"github.com/axellelanca/go_loganizer/internal/config"
	"github.com/axellelanca/go_loganizer/internal/reporter"
	"github.com/spf13/cobra"
)

// Variables pour stocker les valeurs des drapeaux
var (
	configPath string
	outputPath string
)

// analyzeCmd représente la commande analyze
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les fichiers de logs spécifiés dans la configuration",
	Long: `La commande analyze lit un fichier de configuration JSON contenant 
les logs à analyser, puis traite chaque log en parallèle en utilisant 
des goroutines.

Exemple:
  loganalyzer analyze -c config.json -o report.json`,
	Run: runAnalyze, // La fonction qui sera exécutée quand on appelle cette commande
}

func init() {
	// Ajouter la commande analyze à la commande racine
	rootCmd.AddCommand(analyzeCmd)

	// Définir les drapeaux pour cette commande
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Chemin vers le fichier de configuration JSON (obligatoire)")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Chemin vers le fichier de sortie JSON (optionnel)")
	
	// Rendre le drapeau config obligatoire
	analyzeCmd.MarkFlagRequired("config")
}

// runAnalyze contient la logique principale de la commande analyze
func runAnalyze(cmd *cobra.Command, args []string) {
	// 1. Charger la configuration
	fmt.Printf("Chargement de la configuration depuis: %s\n", configPath)
	
	configs, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Erreur lors du chargement de la configuration: %v\n", err)
		return
	}

	fmt.Printf("Configuration chargée: %d logs à analyser\n", len(configs))

	// 2. Préparer la concurrence
	// WaitGroup permet d'attendre que toutes les goroutines se terminent
	var wg sync.WaitGroup
	
	// Canal pour collecter les résultats de manière sécurisée
	// Le canal a une taille égale au nombre de logs pour éviter les blocages
	resultsChan := make(chan analyzer.LogResult, len(configs))

	// 3. Lancer une goroutine pour chaque log
	for _, logConfig := range configs {
		// Incrémenter le compteur de WaitGroup avant de lancer la goroutine
		wg.Add(1)
		
		// Lancer la goroutine (fonction anonyme)
		// On passe logConfig en paramètre pour éviter les problèmes de closure
		go func(lc config.LogConfig) {
			// Décrémenter le compteur quand cette goroutine se termine
			defer wg.Done()
			
			fmt.Printf("Début de l'analyse: %s (%s)\n", lc.ID, lc.Path)
			
			// Analyser le log
			result := analyzer.AnalyzeLog(lc.ID, lc.Path)
			
			// Envoyer le résultat dans le canal
			resultsChan <- result
			
			// Afficher le résultat sur la console
			if result.Status == "OK" {
				fmt.Printf("✓ %s: %s\n", result.LogID, result.Message)
			} else {
				fmt.Printf("✗ %s: %s - %s\n", result.LogID, result.Message, result.ErrorDetails)
			}
		}(logConfig) // Passer logConfig comme argument à la fonction anonyme
	}

	// 4. Attendre que toutes les goroutines se terminent
	wg.Wait()
	
	// Fermer le canal car on n'enverra plus de résultats
	close(resultsChan)

	// 5. Collecter tous les résultats depuis le canal
	var results []analyzer.LogResult
	for result := range resultsChan {
		results = append(results, result)
	}

	// 6. Afficher le résumé
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

	// 7. Exporter vers JSON si demandé
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