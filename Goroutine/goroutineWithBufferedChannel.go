package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 4
	ch <- 5

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// here buffer capacity is 2. If I go to send 3rd value then it will create a deadlock cause there is no receiver for that.
