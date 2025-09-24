package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd représente la commande de base quand elle est appelée sans sous-commandes
var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Un outil d'analyse de logs en parallèle",
	Long: `LogAnalyzer est un outil CLI qui permet d'analyser plusieurs fichiers 
de logs en parallèle et de générer des rapports au format JSON.

Exemple d'utilisation:
  loganalyzer analyze -c config.json -o report.json`,
}

// Execute ajoute toutes les commandes enfant à la commande racine et configure les drapeaux de façon appropriée.
// C'est appelé par main.main(). Elle ne prend qu'un seul paramètre, args.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Ici on peut ajouter des drapeaux persistants et de la configuration
	// Les drapeaux persistants seront disponibles pour cette commande et toutes les sous-commandes

	// Exemple: rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.loganalyzer.yaml)")
}