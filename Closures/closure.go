package main

import "fmt"

// a closure is a function that captures and remembers variables from its surrounding scope,
// even after the scope has exited. Closures are useful for maintaining state and creating function factories.

// Just remember for now. And figure out the mechanism behind.
// Use case: State management -- State transfer across function.
func counter() func() int {

	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {

	increament := counter()

	fmt.Println(increament())
	fmt.Println(increament())
	fmt.Println(increament())
}
