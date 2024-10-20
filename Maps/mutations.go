package main

import (
	"errors"
	"fmt"
)

// INSERT
// GET
// DELETE
// IF EXISTS

type user struct {
	name string
	age  int
}

func deleteIfExists(users map[string]user, name string) (deleted bool, err error) {
	if _, ok := users[name]; !ok {
		return false, errors.New("user not found")
	}
	delete(users, name)
	return true, nil
}

func main() {
	users := map[string]user{
		"AFID": {
			name: "Afid",
			age:  23},

		"AZWAD": {
			name: "Azwad",
			age:  24},
	}

	arr := [3]string{"AFID", "Afid", "AZWAD"}

	for _, name := range arr {
		deleted, err := deleteIfExists(users, name)
		if err != nil {
			fmt.Println(err)
		}
		if deleted {
			fmt.Println("Deleted")
		}
	}
}

// Map keys may be of any type that is comparable.
// keys are not case-sensitive and that's why its deleted "AFID".

// slices, maps and functions are not be compared using '=='.
