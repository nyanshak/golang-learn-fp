package main

import (
	"fmt"
	"sync"
)

var integers = []int{}

func addToSlice(i int, wg *sync.WaitGroup) {
	integers = append(integers, i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	numbersToAdd := 10
	wg.Add(numbersToAdd)

	for i := range numbersToAdd {
		go addToSlice(i, &wg)
	}

	wg.Wait()

	fmt.Println(integers)
}
