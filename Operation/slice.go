package main

import "fmt"

func sliceUse() {

	myslice1 := []int{1, 2, 3, 4, 5, 6}

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myarr1 := [5]int{1, 2, 3, 4, 5}
	myslice2 := myarr1[1:4]

	fmt.Printf("myslice2 = %v\n", myslice2)
}
