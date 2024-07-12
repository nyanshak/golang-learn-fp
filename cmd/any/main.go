package main

import (
	"fmt"
)

func Any[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func All[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func IsEven(n int) bool {
	return n%2 == 0
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Numbers:", numbers)
	fmt.Println("Are any of the numbers even?", Any(numbers, IsEven))
	fmt.Println("Are all of the numbers even?", All(numbers, IsEven))
}
