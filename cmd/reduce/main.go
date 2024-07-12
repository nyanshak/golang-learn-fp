package main

import (
	"fmt"
)

type ReduceFunc[A any] func(a1, a2 A) A

func Reduce[A any](input []A, reducer ReduceFunc[A]) A {
	if len(input) == 0 {
		return *new(A)
	}

	result := input[0]
	for _, elem := range input[1:] {
		result = reducer(result, elem)
	}
	return result
}

func Add(a1, a2 int) int {
	return a1 + a2
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := Reduce(numbers, Add)
	product := Reduce(numbers, func(a1, a2 int) int {
		return a1 * a2
	})

	fmt.Println("sum of    : ", numbers, " is: ", sum)
	fmt.Println("product of: ", numbers, " is: ", product)
}
