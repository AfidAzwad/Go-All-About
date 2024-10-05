package main

import "fmt"

func addNumbers(slice []int, numbers ...int) []int {
	slice = append(slice, numbers...)
	return slice
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main() {
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Total : ", total) // Output: 15

	// Passing a slice to the variadic function
	nums := []int{6, 7, 8, 9, 10}
	fmt.Println("Total of slice: ", sum(nums...)) // Output: 40

	// Append a slice using the variadic operator
	extraNums := []int{11, 12, 13}
	extendedSlice := addNumbers(nums, extraNums...)
	fmt.Println("Extended Slice: ", extendedSlice) // Output: [6 7 8 9 10 11 12 13]
}
