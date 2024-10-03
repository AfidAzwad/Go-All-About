package main

import "fmt"

type Rectangle struct {
	height int
	width  int
}

func (r Rectangle) area() int {
	return r.height * r.width
}

func main() {
	r := Rectangle{
		height: 10,
		width:  20,
	}
	fmt.Println("Area of Rectangle : ", r.area())
}

// here area method takes special parameter of struct. In this pattern we can declare methods for structs.
