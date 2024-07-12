package main

import (
	"fmt"
)

func IsEven(n int) bool {
	return n%2 == 0
}

func IsOdd(n int) bool {
	return n%2 != 0
}

func CheckMatch(n int, f func(int) bool) {
	if f(n) {
		fmt.Printf("%d matches condition\n", n)
		return
	}
	fmt.Printf("%d does not match!\n", n)
}

func main() {
	fmt.Println("IsEven")
	CheckMatch(2, IsEven)

	fmt.Println("IsOdd")
	CheckMatch(2, IsOdd)
}
