package main

import (
	"fmt"
)

func GetGreetFunc(name string) func() {
	if name == "" {
		name = "World"
	}

	return func() {
		fmt.Println("Hello, " + name + "!")
	}
}

func main() {
	// Call function immediately without assigning to a variable.
	GetGreetFunc("")()
	GetGreetFunc("Gopher")()

	// Assign the function to a variable.
	learnGreetFunc := GetGreetFunc("FP learners")

	// Call the function we returned.
	learnGreetFunc()
}
