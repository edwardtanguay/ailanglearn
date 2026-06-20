package helper

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// Category represents a grouping of outline items (e.g., fruits, vegetables).
type Category struct {
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

// GetCLIInfo returns descriptive information about this Go CLI.
func GetCLIInfo() string {
	return "AILangLearn Go CLI - Version 1.0.0\n" +
		"A suite of helper tools for parsing outline data and managing language learning content."
}

// ResolvePath resolves a path relative to the project root directory (by checking for package.json).
func ResolvePath(pathFromRoot string) string {
	// If package.json is in the current directory, we are running from root
	if _, err := os.Stat("package.json"); err == nil {
		return pathFromRoot
	}
	// If package.json is in parent, we are running from one level deep (e.g., inside cli/)
	if _, err := os.Stat("../package.json"); err == nil {
		return filepath.Join("..", pathFromRoot)
	}
	// Fallback to the original path
	return pathFromRoot
}

// ParseOutline parses the outline text file into a slice of Category objects.
func ParseOutline(filePath string) ([]Category, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var categories []Category
	var currentCategory *Category

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		// A category or item starts with "- " after trimming
		if strings.HasPrefix(trimmed, "- ") {
			// Check if the original line had leading whitespace (e.g. tab or spaces)
			hasLeadingWhitespace := len(line) > 0 && (line[0] == '\t' || line[0] == ' ')
			if !hasLeadingWhitespace {
				name := strings.TrimPrefix(trimmed, "- ")
				categories = append(categories, Category{
					Name:  strings.TrimSpace(name),
					Items: []string{},
				})
				currentCategory = &categories[len(categories)-1]
			} else {
				if currentCategory != nil {
					item := strings.TrimPrefix(trimmed, "- ")
					currentCategory.Items = append(currentCategory.Items, strings.TrimSpace(item))
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
