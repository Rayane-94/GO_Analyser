package main

import (
	"fmt"
	"os"

	"github.com/axellelanca/go_loganizer/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}