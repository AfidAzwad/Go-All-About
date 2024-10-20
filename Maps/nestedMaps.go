package main

import "fmt"

func main() {
	nestedMap := make(map[string]map[string]int)

	// Assign values using map literals
	nestedMap["User1"] = map[string]int{
		"Age": 22,
	}

	nestedMap["User2"] = map[string]int{
		"Age": 25,
	}

	fmt.Println(nestedMap) //map[User1:map[Age:22] User2:map[Age:25]]

	nestedMap2 := map[rune]map[string]int{
		'3': {
			"Age": 22,
		},
		'4': {
			"Age": 25,
		},
	}
	fmt.Println(nestedMap2) //map[51:map[Age:22] 52:map[Age:25]]
}

// rune is a type in Go that represents a Unicode code point and is an alias for int32.
// It allows for the representation of characters from various languages and scripts,
// enabling text processing that is agnostic to the specific language.
// For example, when processing text, we can use rune to iterate over characters
// and handle Unicode characters properly, ensuring accurate character manipulation
// regardless of the language.
