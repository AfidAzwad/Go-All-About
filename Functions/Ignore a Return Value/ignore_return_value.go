package main

import "fmt"

func ignoreReturn() (int, int) {
	return 3, 4
}
func main() {
	x, _ := ignoreReturn()
	fmt.Println(x)
}

// if we don't need y value then we need to ignore it cause GO throw error if we don't use any variable
