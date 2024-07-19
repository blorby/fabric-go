package config

import (
    "fmt"
    "os"
    "path/filepath"
)

var (
    HomeDir    string
    ConfigDir  string
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
    // Implement setup logic here
    fmt.Println("Setting up Fabric...")
    // Create necessary directories, download patterns, etc.
}
