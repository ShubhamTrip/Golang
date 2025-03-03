package main

import "fmt"

func arrayUse() {

	var arr1 = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Array 1: ", arr1)

	arr2 := [...]string{"apple", "mango", "banana", "grapes"}

	fmt.Println("Array 2: ", arr2)

	arr3 := [...]int{3: 10, 5: 50}

	fmt.Println("Array 3: ", arr3)

}
