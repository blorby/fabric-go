package patterns

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/blorby/fabric-go/internal/config"
)

func ListPatterns() {
	entries, err := os.ReadDir(config.PatternsDir)
	if err != nil {
		fmt.Printf("Error reading patterns directory: %v\n", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}
}

func RunPattern(patternName, input string) (string, error) {
	systemFile := filepath.Join(config.PatternsDir, patternName, "system.md")
	userFile := filepath.Join(config.PatternsDir, patternName, "user.md")

	systemContent, err := os.ReadFile(systemFile)
	if err != nil {
		return "", fmt.Errorf("error reading system file: %v", err)
	}

	userContent, err := os.ReadFile(userFile)
	if err != nil {
		return "", fmt.Errorf("error reading user file: %v", err)
	}

	// Simple placeholder implementation
	// In a real implementation, you would use these contents with an AI service
	result := fmt.Sprintf("Pattern: %s\nSystem: %s\nUser: %s\nInput: %s",
		patternName,
		strings.TrimSpace(string(systemContent)),
		strings.TrimSpace(string(userContent)),
		input)

	return result, nil
}
