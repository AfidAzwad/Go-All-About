package main

import "fmt"

type Waiter struct {
	name string
	Table
}
type Table struct {
	number int
	seater int
	shape  string
}

func server(w Waiter) {
	fmt.Println("Waiter Name :", w.name)
	fmt.Println("Assigned Table :", w.number)
	fmt.Println("Table Capacity", w.seater)
	fmt.Println("Table Shape", w.shape)
	fmt.Println("............................")
}
func main() {
	server(Waiter{
		name: "Afid",
		Table: Table{
			number: 5,
			seater: 2,
			shape:  "square",
		},
	})
	server(Waiter{
		name: "Azwad",
		Table: Table{
			number: 2,
			seater: 4,
			shape:  "square",
		},
	})
}

// Go does not support classical inheritance. Instead, it promotes composition through embedding.
// Here Table struct embedded in Waiter struct.

/*Go encourages composition over inheritance and provides a more lightweight, flexible form of
polymorphism through interfaces, allowing Go to use certain OOP patterns while staying simpler and more efficient. */
