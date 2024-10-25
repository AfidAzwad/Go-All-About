package main

import (
	"fmt"
	"os"
)

func writeFile(filename string, content string) {
	// Create a file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Ensure the file is closed after function completes
	defer fmt.Println("File closed successfully!!")
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File written successfully")
}

func main() {
	writeFile("example.txt", "Hello, World!")
}

/* Since defer statements execute in LIFO (Last In, First Out) order, here’s what happens:

1. First, defer fmt.Println("File closed successfully!!") is added to the stack.
2. Then, defer file.Close() is added on top of it.

When the function exits, deferred statements execute in reverse order, meaning file.Close()
will run before fmt.Println("File closed successfully!!"). */
