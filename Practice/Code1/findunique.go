package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter the number of values")

	var n int

	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		isUnique := true
		for j := 0; j < n; j++ {
			if arr[i] == arr[j] && i != j {
				isUnique = false
			}
		}
		if isUnique == true {
			fmt.Println(arr[i], "is Unique")
		}
	}
}
