package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	HomeDir     string
	ConfigDir   string
	PatternsDir string
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}

	HomeDir = homeDir
	ConfigDir = filepath.Join(HomeDir, ".config", "fabric")
	PatternsDir = filepath.Join(ConfigDir, "patterns")
}

func Setup() {
	fmt.Println("Setting up Fabric...")

	if err := createDirectories(); err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		os.Exit(1)
	}

	// Here you would add logic to download or copy initial patterns
	fmt.Println("Setup complete. Patterns directory created at:", PatternsDir)
}

func createDirectories() error {
	dirs := []string{ConfigDir, PatternsDir}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	return nil
}
