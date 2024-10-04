package main

import (
	"fmt"
)

// Function to identify type and print based on type
func identifyType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("This is an int: %d\n", v)
	case string:
		fmt.Printf("This is a string: %s\n", v)
	case float64:
		fmt.Printf("This is a float64: %f\n", v)
	default:
		fmt.Println("Unknown type.")
	}
}

func main() {
	identifyType(42)         // int
	identifyType("Hello Go") // string
	identifyType(3.14)       // float64
	identifyType(true)       // default case, unknown type
}

/* Here if we don't use interface then we need to declare what type of
data that method will receive as parameter. So here interface made it generalized. */
