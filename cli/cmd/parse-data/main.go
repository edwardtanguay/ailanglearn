package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"ailanglearn/cli/pkg/helper"
)

func main() {
	// 1. Resolve paths dynamically relative to workspace root
	inputPath := helper.ResolvePath("data/test.outline.dpod.txt")
	outputPath := helper.ResolvePath("data-parsed/test.json")

	fmt.Printf("Reading outline from: %s\n", inputPath)

	// 2. Parse the outline file using helper package
	categories, err := helper.ParseOutline(inputPath)
	if err != nil {
		log.Fatalf("Error parsing outline file: %v\n", err)
	}

	// 3. Serialize to JSON with formatting
	jsonData, err := json.MarshalIndent(categories, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling outline to JSON: %v\n", err)
	}

	// 4. Ensure directory of output path exists
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Error creating output directory: %v\n", err)
	}

	// 5. Write file
	if err := os.WriteFile(outputPath, jsonData, 0644); err != nil {
		log.Fatalf("Error writing JSON output file: %v\n", err)
	}

	// 6. Display success info
	fmt.Printf("Successfully parsed %d categories to: %s\n", len(categories), outputPath)
	for _, cat := range categories {
		fmt.Printf(" - %s (%d items)\n", cat.Name, len(cat.Items))
	}
}
