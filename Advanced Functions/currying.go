package main

import "fmt"

func addition(a, b int) int {
	return a + b
}

func multiplication(a, b int) int {
	return a * b
}

func setOperation(a int, arithmetic func(int, int) int) func(int) int {
	return func(b int) int {
		return arithmetic(a, b)
	}
}

func main() {
	additionOperation := setOperation(10, addition)
	multiplicationOperation := setOperation(10, multiplication)

	additionResult := additionOperation(20)
	mulResult := multiplicationOperation(20)

	fmt.Println("Addition : ", additionResult)
	fmt.Println("Multiplication : ", mulResult)
}

/* Function currying is the practice of writing a function that takes a function(or functions)
   as input and returns a new function. */

/* Here func(b int) int {} is an anonymous function because it has no name. We need this cause if directly call 'return arithmetic(a, b)'
it will return an integer but setOperation function is set to return a function */
