package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func get(s Shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func main() {
	r := Rectangle{12, 34}
	c := Circle{10}

	fmt.Println("Rectangle =>")
	get(r)
	fmt.Println("----------------------------")
	fmt.Println("Circle =>")
	get(c)
}

/* Interfaces are collections of method signatures. A type "implements" an interface if it has all
of the methods of the given interface defined on it. */
