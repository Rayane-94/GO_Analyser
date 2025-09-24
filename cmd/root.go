package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Un outil d'analyse de logs en parallèle",
	Long: `LogAnalyzer est un outil CLI qui permet d'analyser plusieurs fichiers 
de logs en parallèle et de générer des rapports au format JSON.

Exemple d'utilisation:
  loganalyzer analyze -c config.json -o report.json`,
}

func Execute() error {
	return rootCmd.Execute()
}