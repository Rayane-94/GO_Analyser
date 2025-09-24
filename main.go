package main

import (
	"fmt"
	"os"

	"github.com/axellelanca/go_loganizer/cmd"
)

// La fonction main est le point d'entrée de notre application
func main() {
	// Si on a une erreur lors de l'exécution de la commande racine, on l'affiche et on quitte
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}