package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Anwendung: flow [.flow Datei]")
		os.Exit(65)
	}

	if filepath.Ext(os.Args[1]) != ".flow" {
		fmt.Println("Akzeptiert werden nur .flow Dateien")
		os.Exit(67)
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Fehler beim Lesen der Datei", err)
		os.Exit(66)
	}

	input := string(bytes)

	tokens := lexer(input)
	fmt.Println(tokens)
}
