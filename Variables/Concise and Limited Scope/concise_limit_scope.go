package main

import "fmt"

func getLength(email string) int {
	return len(email)
}

func main() {
	email := "example@example.com"
	if length := getLength(email); length < 1 {
		fmt.Printf("The length of your email is %v and Email is invalid!\n", length)
	} else {
		fmt.Printf("The length of your email is %v and Email is valid!\n", length)
	}
}

// here length is not available to access from parent scope. So we made limited scope.
