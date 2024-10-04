package main

import "fmt"

func main() {
	// Basic for loop : Initialization, Condition and Post (increment or decrement)
	fmt.Println("Basic for loop ===> ")

	for i := 0; i < 10; i++ {
		fmt.Println("Iteration", i)
	}

	// Go doesn't have a while loop, but you can mimic a while loop by using a for loop with only a condition.
	fmt.Println("for loop like while ===> ")

	i := 0
	for i < 5 {
		fmt.Println("Iteration:", i)
		i++
	}

	// Infinite for loop - if we don't use break then it will not stop
	fmt.Println("Infinite for loop ===> ")

	index := 0
	for {
		fmt.Println("Iteration:", i)
		index++
		if index == 5 {
			break
		}
	}

	// for loop with range
	fmt.Println("for loop with range ===> ")

	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
