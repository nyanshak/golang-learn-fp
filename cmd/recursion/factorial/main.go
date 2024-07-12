package main

import "fmt"

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println("factorial( 3) = ", factorial(3))
	fmt.Println("factorial( 5) = ", factorial(5))
	fmt.Println("factorial(10) = ", factorial(10))
}
