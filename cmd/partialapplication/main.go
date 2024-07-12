package main

import (
	"fmt"
)

func Greet(greeting, name string) string {
	return fmt.Sprintf("%s, %s", greeting, name)
}

func PrefixGreet(greeting string) func(string) string {
	return func(name string) string {
		return Greet(greeting, name)
	}
}

func main() {
	gdayFn := PrefixGreet("G'day")
	yoFn := PrefixGreet("Yo")

	fmt.Println(gdayFn("Gophers"))
	fmt.Println(yoFn("class"))
}
