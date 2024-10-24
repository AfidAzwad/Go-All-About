package main

import (
	"fmt"
	"time"
)

func sendOnly(ch chan<- int) {
	ch <- 42 // Can only send
}

func receiveOnly(ch <-chan int) {
	fmt.Println(<-ch) // Can only receive
	fmt.Println("Printed.....")
}

func main() {
	ch := make(chan int)

	go sendOnly(ch)
	go receiveOnly(ch)

	//fmt.Println(<-ch)
	time.Sleep(10 * time.Second)
}

/* Here I have commented out a Println line in the main function cause for this receiveOnly does not get value.

Why It's Not Working:
 	1. In the main function, the line fmt.Println(<-ch) blocks and waits for data from the channel.
	2. Since the main function itself is trying to receive from the channel, the receiveOnly goroutine
		never gets a chance to read from the channel (because only one receiver can read from the channel at a time).
	3. When fmt.Println(<-ch) in main receives the value, it removes the value from the channel,
		so there's nothing left for the receiveOnly function to read from. */
