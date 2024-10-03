package main

import "fmt"

type Mover interface {
	move() string
}

type Sayer interface {
	say() string
}

type Car struct {
	brand string
}

// Implementing the move() method for the Car struct
func (c Car) move() string {
	return c.brand + " car is moving."
}

type Person struct {
	name string
}

// Implementing the say() method for the Person struct
func (p Person) say() string {
	return p.name + " says: Hello!"
}

// Function that takes both a Mover and a Sayer as parameters
func action(m Mover, s Sayer) {
	fmt.Println(m.move())
	fmt.Println(s.say())
}

func main() {
	// Create instances of Car and Person
	car := Car{brand: "Toyota"}
	person := Person{name: "Azwad"}
	action(car, person)
}
