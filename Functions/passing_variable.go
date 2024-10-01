package main

import "fmt"

func increment(x int) {
	x++
}
func main() {
	x := 5
	increment(x)
	fmt.Println(x)
}

// output is 5 cause the increment function received a copy of x
// variables in Go are passed by value (except for a few data types).
// "Pass by value" means that when a variable is passed into a function, that function receives a copy
// of variable. The function unable to mutate the caller's data.
