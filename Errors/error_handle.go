package main

import (
	"fmt"
	"os"
)

func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	// Close the file after the function completes
	defer file.Close()
	fmt.Println("File opened successfully:", file.Name())
	return nil
}

func main() {
	// Try to open a non-existent file to trigger an error
	err := readFile("nonexistent.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// If no error, continue with normal execution
	fmt.Println("File read successfully!")
}

// Python raises exception types but in Go, an error is just another value that we handle like any other value.
