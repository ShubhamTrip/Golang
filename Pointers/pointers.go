package main

type Circle struct {
	radius float64
	area   float64
}

func (c *Circle) calculateArea() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	circle := Circle{radius: 5, area: 0}
	circle.calculateArea()
	println("Area of circle:", circle.area)
	// Output: Area of circle: 78.5
}
