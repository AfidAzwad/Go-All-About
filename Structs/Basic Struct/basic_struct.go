package main

import "fmt"

type Wheel struct {
	Radius   int
	Material string
}

type Car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	BackWheel  Wheel
}

func canCarBuild(myCar Car) bool {
	if myCar.Height == 0 {
		return false
	}
	if myCar.Width == 0 {
		return false
	}
	if myCar.FrontWheel.Radius == 0 {
		return true
	}
	if myCar.BackWheel.Radius == 0 {
		return false
	}
	return true
}
func main() {
	myCar := Car{}
	myCar.FrontWheel.Radius = 5
	myCar.FrontWheel.Material = "White"
	myCar.BackWheel.Radius = 6
	myCar.BackWheel.Material = "Black"
	fmt.Println(myCar)
	fmt.Println(canCarBuild(myCar))
}

// Go has no concept of classes or inheritance. Instead, Go uses structs to define data structures.
