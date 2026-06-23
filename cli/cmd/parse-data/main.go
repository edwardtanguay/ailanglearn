package main

import (
	"fmt"
	"log"

	"ailanglearn/cli/pkg/helper"
	"ailanglearn/cli/pkg/parser"
)

func main() {
	// 1. Resolve paths dynamically relative to workspace root
	inputPath := helper.ResolvePath("data/test.outline.dpod.txt")
	outputPath := helper.ResolvePath("data-parsed/test.json")

	fmt.Printf("Reading outline from: %s\n", inputPath)
	fmt.Printf("Writing JSON outline to: %s\n", outputPath)

	// 2. Instantiate OutlineParser (this runs the parsing pipeline internally)
	_, err := parser.New(inputPath, outputPath)
	if err != nil {
		log.Fatalf("Error parsing outline: %v\n", err)
	}

	fmt.Println("Successfully parsed outline!")
}
