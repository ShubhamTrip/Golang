package main

import (
	"fmt"
	"math"
)

// Define the interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle type that implements the Shape interface
type Circle struct {
	radius float64
}

// Rectangle type that implements the Shape interface
type Rectangle struct {
	length, width float64
}

// Here since we have added Interface above, we need to implement the methods
// If we don't implement the methods, we will get an error
// Here dont have to use implements keyword, it is automatically implemented

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Main function to demonstrate the interface
func main() {
	var s Shape

	s = Circle{radius: 5}
	fmt.Println("C Area:", s.Area())
	fmt.Println("C Perimeter:", s.Perimeter())

	s = Rectangle{length: 4, width: 3}
	fmt.Println("R Area:", s.Area())
	fmt.Println("R Perimeter:", s.Perimeter())
}
