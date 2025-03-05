package main

import (
	"fmt"
	"sort"
)

func sliceUse() {

	myslice1 := []int{1, 2, 3, 4, 5, 6}

	myslice1 = append(myslice1, 20, 21)
	// capacity of slice is doubled when it reaches its limit
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myarr1 := [5]int{1, 2, 3, 4, 5}
	myslice2 := myarr1[1:4]

	fmt.Printf("myslice2 = %v\n", myslice2)

	myslice3 := make([]int, 5) // length = 5, capacity = 5
	// myslice3 := make([]int, 5, 10) // length = 5, capacity = 10
	// myslice3 := make([]int, 0, 10) // length = 0, capacity = 10
	// we can fill more than 5 elements in this slice as capacity is 5
	// but we can fill more elements using append function
	// myslice3 = append(myslice3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// fmt.Printf("myslice3 = %v\n", myslice3)
	// Here capacity increases to 10
	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("capacity = %d\n", cap(myslice3))

	myslice4 := []int{1, 8, 7, 3, 5}

	sort.Ints(myslice4)
	fmt.Println("Sorted slice: ", myslice4)

	fmt.Println(sort.IntsAreSorted(myslice4))

	// How to remove an element from a slice

	myslice5 := []string{"apple", "mango", "banana", "grapes"}

	// remove mango from the slice

	myslice5 = append(myslice5[:1], myslice5[2:]...)

	fmt.Println("Slice after removing mango: ", myslice5)

}
