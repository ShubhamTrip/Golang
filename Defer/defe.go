package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Func ends here...")
	}()
	fmt.Println("The sum is:", add(10, 20, 30))
}

func add(num ...int) int {
	sum := 0
	for _, value := range num {
		sum += value
	}
	return sum
}
