package main

import "fmt"

func main() {
	// Define the anonymous struct inside the main function
	myCar := struct {
		name  string
		model string
	}{
		name:  "Audi",
		model: "A15",
	}

	// Print the values of the anonymous struct
	fmt.Println("Car Name:", myCar.name)
	fmt.Println("Car Model:", myCar.model)
}

// anonymous struct in Go is like a normal struct, but it is defined without a name and therefore can't be referenced elsewhere in the code.
