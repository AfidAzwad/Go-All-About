package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address // Nested struct
}

func main() {
	p := Person{
		Name: "Azwad",
		Age:  30,
		Address: Address{
			City:    "Dhaka",
			Country: "BD",
		},
	}

	// Access nested fields
	fmt.Println("Name:", p.Name)
	fmt.Println("City:", p.Address.City)
	fmt.Println("Country:", p.Address.Country)
}

// In nested structs, one struct is defined inside another, and you must reference the inner struct explicitly to access its fields.
