package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(i) // This will only print odd numbers
	}
}
