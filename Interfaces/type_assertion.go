package main

import "fmt"

// interface
type Animal interface {
	Speak() string
}

// types Fish and Bird implementing the Animal interface
type Fish struct{}
type Bird struct{}

// Speak method for Fish
func (f Fish) Speak() string {
	return "Blub Blub"
}

// Speak method for Bird
func (b Bird) Speak() string {
	return "Chirp Chirp"
}

// Function to check the type of the animal and assert it as Fish or Bird
func identifyAnimal(a Animal) {
	fish, ok := a.(Fish)
	if ok {
		fmt.Println("Fish speaks like :", fish.Speak())
	}
	bird, ok := a.(Bird)
	if ok {
		fmt.Println("Bird speaks like :", bird.Speak())
	}
}

func main() {
	// Create instances of Fish and Bird
	fish := Fish{}
	bird := Bird{}

	// Call identifyAnimal with different types
	identifyAnimal(fish)
	identifyAnimal(bird)
}
