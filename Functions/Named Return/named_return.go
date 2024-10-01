package main

import "fmt"

// if we declare x and y in return info then we need to declare and return the variables.
func namedReturn() (x, y int) {
	return
}

// if we don't declare x and y in return info then we need to declare and return the variables.
func namedReturn2() (int, int) {
	var x int
	var y int
	return x, y
}

func main() {
	fmt.Println(namedReturn())
	fmt.Println(namedReturn2())
}

// output : 0, 0
//          0, 0

// both function will return x and y with 0.
