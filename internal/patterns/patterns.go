package patterns

import (
    "fmt"
    "io/ioutil"
    "path/filepath"

    "github.com/yourusername/fabric-go/internal/config"
)

func ListPatterns() {
    patterns, err := ioutil.ReadDir(config.PatternsDir)
    if err != nil {
        fmt.Printf("Error reading patterns directory: %v\n", err)
        return
    }

    for _, pattern := range patterns {
        if pattern.IsDir() {
            fmt.Println(pattern.Name())
        }
    }
}

func RunPattern(patternName, input string) (string, error) {
    systemFile := filepath.Join(config.PatternsDir, patternName, "system.md")
    userFile := filepath.Join(config.PatternsDir, patternName, "user.md")

    systemContent, err := ioutil.ReadFile(systemFile)
    if err != nil {
        return "", fmt.Errorf("error reading system file: %v", err)
    }

    userContent, err := ioutil.ReadFile(userFile)
    if err != nil {
        return "", fmt.Errorf("error reading user file: %v", err)
    }

    // Here you would implement the logic to process the input using the pattern
    // This might involve calling an AI service or implementing the logic directly

    return "Processed result", nil
}
