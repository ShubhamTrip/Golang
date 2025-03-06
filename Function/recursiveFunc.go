package main

import "fmt"

func countupto10(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return countupto10(x + 1)
}

// Recursive factorial function

func factorial(n int) (result int) {
	if n == 0 {
		return 1
	}
	result = n * factorial(n-1)

	return result
}
