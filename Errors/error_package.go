package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("something went wrong") // Create a new error
	}
	return a / b, nil
}

func main() {
	a := 10
	b := 0
	// Attempt to divide by zero
	result, err := divide(a, b)

	fmt.Printf("A = %v and B = %v == >\n", a, b)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	c := 10
	d := 2
	fmt.Printf("A = %v and B = %v == >\n", c, d)
	result, err = divide(c, 2)
	if err != nil {
		fmt.Println("Error:", err) // Handle the error
	} else {
		fmt.Println("Result:", result) // Print the result if no error
	}
}
