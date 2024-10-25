package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func aggregate(a, b, c int, operation func(int, int) int) int {
	return operation(operation(a, b), c)
}

func main() {
	// Passing the 'add' function to 'aggregate'
	sum := aggregate(5, 6, 7, add)
	fmt.Println(sum) // Output: 18

	// Passing the 'mul' function to 'aggregate'
	product := aggregate(5, 6, 7, mul)
	fmt.Println(product) // Output: 210
}

// here arithmetic is a variable not any specific function type
