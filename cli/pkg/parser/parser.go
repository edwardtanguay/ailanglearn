package parser

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// LineObject represents a single line's indentation level and raw text.
type LineObject struct {
	Level   int    `json:"level"`
	RawText string `json:"raw_text"`
}

// OutlineParser encapsulates the input path, output path, parsed lines, and LineObjects.
type OutlineParser struct {
	inputPath   string
	outputPath  string
	lines       []string
	lineObjects []LineObject
}

// New constructs an OutlineParser and executes the parsing steps synchronously.
func New(inputPath, outputPath string) (*OutlineParser, error) {
	p := &OutlineParser{
		inputPath:  inputPath,
		outputPath: outputPath,
	}

	if err := p.parseIntoLines(); err != nil {
		return nil, err
	}

	if err := p.createLineObjects(); err != nil {
		return nil, err
	}

	if err := p.createJsonFile(); err != nil {
		return nil, err
	}

	return p, nil
}

// parseIntoLines reads lines from the input file and stores them.
func (p *OutlineParser) parseIntoLines() error {
	file, err := os.Open(p.inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p.lines = append(p.lines, scanner.Text())
	}
	return scanner.Err()
}

// createLineObjects processes the lines to compute level and raw text.
func (p *OutlineParser) createLineObjects() error {
	for _, line := range p.lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		level := calculateLevel(line)

		p.lineObjects = append(p.lineObjects, LineObject{
			Level:   level,
			RawText: trimmed,
		})
	}
	return nil
}

// createJsonFile converts the lineObjects into a formatted JSON file.
func (p *OutlineParser) createJsonFile() error {
	jsonData, err := json.MarshalIndent(p.lineObjects, "", "  ")
	if err != nil {
		return err
	}

	outputDir := filepath.Dir(p.outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	return os.WriteFile(p.outputPath, jsonData, 0644)
}

// calculateLevel calculates the nesting level from leading whitespace.
func calculateLevel(line string) int {
	trimmed := strings.TrimLeft(line, "\t ")
	indentStr := line[:len(line)-len(trimmed)]

	tabs := strings.Count(indentStr, "\t")
	spaces := strings.Count(indentStr, " ")

	// Treat 4 spaces as 1 indent level.
	level := tabs + (spaces / 4) + 1
	return level
}
