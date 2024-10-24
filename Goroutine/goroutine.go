package main

import (
	"fmt"
	"sync"
)

func addition(a, b int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes.
	fmt.Println("Addition Started ==== ")
	fmt.Println("a + b =", a+b)
	fmt.Println("Addition Ended ========== ")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // Add a goroutine to the WaitGroup
	go addition(5, 6, &wg)
	//time.Sleep(2 * time.Second) // 2 seconds break
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Hello World !")
}

// without WaitGroup main function will not wait for addition function.
