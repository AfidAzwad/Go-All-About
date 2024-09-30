package main

import "fmt"

func main() {
	// initialize variables here
	var limit int
	var cost float64
	var canBuy bool
	var productName string

	fmt.Printf("%v\n%f\n%v\n%s\n", limit, cost, canBuy, productName)
}

// Printf used to print formatted string
// %v is for default representation
// %s Interpolate a string
// %d Interpolate an integer in decimal form
// %f interpolate a decimal
