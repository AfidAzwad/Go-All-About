package main

import "fmt"

func Addition(a, b int, ch chan int) {
	result := a + b
	ch <- result // it has receiver in the main function which is the last 'Println'
}

func main() {
	ch := make(chan int)

	// Start a goroutine to receive from the channel first
	go func() {
		fmt.Println(<-ch) // Receives 55
	}()

	ch <- 55 // This won't block because the goroutine is ready to receive

	go Addition(5, 6, ch) // Function that sends result to channel

	fmt.Println(<-ch)
}
