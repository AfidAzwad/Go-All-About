package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
}

// map in Go can be used as Hash Table
// Looking up a value in a map by its key is much faster than searching through a slice.
