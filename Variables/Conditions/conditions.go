package main

import "fmt"

func main() {
	height := 6
	weight := 80

	if weight > height {
		fmt.Println("Weight is greater than Height")
	} else {
		fmt.Println("Height is greater than Weight")
	}
}
