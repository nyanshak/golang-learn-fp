package main

import "fmt"

func Filter[T any](originalSlice []T, f func(T) bool) []T {
	filteredSlice := []T{}

	for _, v := range originalSlice {
		if f(v) {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func IsEven(n int) bool {
	return n%2 == 0
}

func GetInRangeFunc(minThreshold, maxThreshold int) func(int) bool {
    return func(n int) bool {
        return n >= minThreshold && n <= maxThreshold
    }
}

func main() {
	originalList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evens := Filter(originalList, IsEven)
	fmt.Println(evens)

    statuses := []int{200, 400, 500, 503, 429, 418}

    status400Filter := GetInRangeFunc(400, 499)
    fmt.Println(Filter(statuses, status400Filter))
}
