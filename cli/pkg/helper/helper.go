package helper

import (
	"os"
	"path/filepath"
)

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
