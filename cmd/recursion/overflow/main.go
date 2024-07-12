// Overflow demonstrates a stack overflow. By setting a larger stack size,
// you can see that the program completes successfully.
package main

import (
	"flag"
	"fmt"
	"runtime/debug"
)

func f(i int) string {
	// This is set to a very large value to cause a high level of recursion
	// and trigger a stack overflow (when stack size is not manually increased
	// beyond the default value).
	if i == 20_000_000 {
		return "completed successfully"
	}
	return f(i + 1)
}

func main() {
	defaultMaxStackSize := debug.SetMaxStack(0)
	var stackSize int

	flag.IntVar(&stackSize, "stack-size", debug.SetMaxStack(defaultMaxStackSize), "stack size")
	flag.Parse()

	if stackSize > 0 {
		sz := debug.SetMaxStack(stackSize)
		fmt.Println("original stack size:", sz)
		fmt.Println("set stack size:", stackSize)
	}

	fmt.Printf("%q", f(0))
}
