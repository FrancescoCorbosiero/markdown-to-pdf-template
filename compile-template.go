package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	//"strings"
)

/**
 * Script per compilare il template WordPress con i dati del cliente
 *
 * Utilizzo:
 * go run compile-template.go -template template.md -data data.json -output output.md
 */

func main() {
	// Definizione dei flag della riga di comando
	templatePath := flag.String("template", "", "Percorso del file template")
	dataPath := flag.String("data", "", "Percorso del file JSON con i dati")
	outputPath := flag.String("output", "", "Percorso del file di output")

	// Parsing degli argomenti
	flag.Parse()

	// Verifica che tutti i parametri necessari siano stati forniti
	if *templatePath == "" || *dataPath == "" || *outputPath == "" {
		fmt.Println("Utilizzo: go run compile-template.go -template template.md -data data.json -output output.md")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Lettura del file template
	templateContent, err := os.ReadFile(*templatePath)
	if err != nil {
		fmt.Printf("Errore: Impossibile leggere il file template - %v\n", err)
		os.Exit(1)
	}

	// Lettura del file JSON
	dataContent, err := os.ReadFile(*dataPath)
	if err != nil {
		fmt.Printf("Errore: Impossibile leggere il file dati - %v\n", err)
		os.Exit(1)
	}

	// Parsing del JSON
	var data map[string]interface{}
	if err := json.Unmarshal(dataContent, &data); err != nil {
		fmt.Println("Errore: Il file JSON contiene errori di sintassi")
		os.Exit(1)
	}

	// Compila il template
	compiledContent := string(templateContent)

	// Sostituisci tutte le variabili
	for key, value := range data {
		// Converti il valore in stringa
		strValue := fmt.Sprintf("%v", value)
		
		// Crea il pattern regex per trovare {{key}}
		regex := regexp.MustCompile(`{{\s*` + regexp.QuoteMeta(key) + `\s*}}`)
		
		// Sostituisci tutte le occorrenze
		compiledContent = regex.ReplaceAllString(compiledContent, strValue)
	}

	// Scrivi il file output
	err = os.WriteFile(*outputPath, []byte(compiledContent), 0644)
	if err != nil {
		fmt.Printf("Errore: Impossibile scrivere il file di output - %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Template compilato con successo: %s\n", *outputPath)
}