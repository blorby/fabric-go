package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/blorby/fabric-go/internal/config"
	"github.com/blorby/fabric-go/internal/patterns"
)

func main() {
	var (
		text         string
		copy         bool
		pattern      string
		setup        bool
		listPatterns bool
	)

	flag.StringVar(&text, "text", "", "Text to process")
	flag.BoolVar(&copy, "copy", false, "Copy the response to clipboard")
	flag.StringVar(&pattern, "pattern", "", "The pattern to use")
	flag.BoolVar(&setup, "setup", false, "Set up your fabric instance")
	flag.BoolVar(&listPatterns, "list", false, "List available patterns")
	flag.Parse()

	if setup {
		config.Setup()
		return
	}

	if listPatterns {
		patterns.ListPatterns()
		return
	}

	if pattern == "" {
		fmt.Println("Error: Pattern is required")
		os.Exit(1)
	}

	result, err := patterns.RunPattern(pattern, text)
	if err != nil {
		fmt.Printf("Error running pattern: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)

	if copy {
		// Implement clipboard functionality
	}
}
