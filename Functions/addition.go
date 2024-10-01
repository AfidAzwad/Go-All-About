package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}
func main() {
	fmt.Println(concat("hello", " Azwad"))
	fmt.Println(concat("How are", " you?"))
}

// we need to mention date type what the function will return. Here it is 'string'
