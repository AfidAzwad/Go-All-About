package main

import "fmt"

func getInsuranceAmount(amount int) int {
	if amount != 5000 {
		return amount + 1
	}
	if amount > 1000 {
		return amount + 100
	}
	return amount
}
func main() {
	fmt.Println(getInsuranceAmount(5000))
}

// here we have used guard clauses. It leverages the ability to return early.
