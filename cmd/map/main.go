package main

import (
	"fmt"
)

func Map[A, B any](s []A, f func(A) B) []B {
	b := make([]B, len(s))

	for idx := range s {
		b[idx] = f(s[idx])
	}

	return b
}

func Square(x int) int {
	return x * x
}

func IsEven(x int) bool {
    return x%2 == 0
}

func main() {
	fmt.Println("Square of [1 2 3]:")
	fmt.Println(Map([]int{1, 2, 3}, Square))

    fmt.Println("IsEven of [1 2 3]:")
    fmt.Println(Map([]int{1, 2, 3}, IsEven))
}
