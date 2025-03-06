package main

import (
	"fmt"
)

// Here will we see Call by value only

// Function are of two types:
// 1. Built-in functions
// 2. User-defined functions

// Built-in functions are those functions which are already defined in the Go language. For example: len(), cap(), append(), make(), new(), copy(), close(), delete(), panic(), recover(), print(), println(), etc.

// User-defined functions are those functions which are defined by the user. For example: main(), sum(), sub(), mul(), div(), etc.

// Main function is the entry point of the program. It is the first function that will be executed.
func main() {

	function1()
	function2(10, 20)
	result := function3()
	fmt.Println("This is function4 and the sum of a and b is: ", result)
	result = function4(10, 20)
	fmt.Println("This is function4 and the sum of a and b is: ", result)
	result1, result2 := function5(10, 20)
	fmt.Println("This is function5 and the sum of a and b is: ", result1)
	fmt.Println("This is function5 and the sub of a and b is: ", result2)
	result1, result2 = function6(10, 20)
	fmt.Println("This is function6 and the sum of a and b is: ", result1)
	fmt.Println("This is function6 and the sub of a and b is: ", result2)
	result = sum(10, 20, 30, 40, 50)
	fmt.Println("This is function7 and the sum of a and b is: ", result)
	result = function7("Shubham", 20, 30, 40, 50)
	fmt.Println("This is function7 and the sum of a and b is: ", result)

	//calling recursive function
	countupto10(1)
	value := factorial(7)

	fmt.Println("Factorial of 7 is: ", value)

}

// Functions are of four types:
// 1. Function with no return type and no arguments

func function1() {

	a := 10
	b := 20

	fmt.Println("This is function1 and the sum of a and b is: ", a+b)
}

// 2. Function with no return type and arguments

func function2(a int, b int) { // function2(a, b int) is also correct

	fmt.Println("This is function2 and the sum of a and b is: ", a+b)
}

// 3. Function with return type and no arguments
func function3() int {

	a := 10
	b := 20

	return a + b
}

// 4. Function with return type and arguments
func function4(a int, b int) int { // function4(a, b int) is also correct

	return a + b
}

// Function with multiple return values
func function5(a int, b int) (int, int) {

	return a + b, a - b
}

// Function with named return values
func function6(a int, b int) (sum int, sub int) {

	sum = a + b
	sub = a - b

	return
}

// Function with variadic parameters
// Here we dont know how many arguments will be passed to the function
// So we use ... before the type of the argument
// We can pass 0 or more arguments to the function
// We can pass the arguments of the same type only
func sum(nums ...int) int {
	total := 0

	// We use _ if we dont want to use the value of the variable. This is called write-only variable
	// Here we are using range to iterate over the slice
	// range returns two values, first is the index and second is the value
	// We are using _ to ignore the index value.
	for _, n := range nums {
		total += n
	}
	return total
}

// Function with variadic parameters
// Here we can pass arguments of different types
func function7(name string, b ...int) int {

	fmt.Println(name)
	fmt.Println(b)

	sum := 0

	for _, v := range b {
		sum += v
	}

	return sum
}
