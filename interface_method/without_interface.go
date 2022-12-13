package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return (c.radius * math.Pi * c.radius)
}
func (c circle) perimeter() float64 {
	return (2 * math.Pi * c.radius)
}
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}
func printCircleDetails(c circle) {
	fmt.Printf(`Area of cirlce is: %f 
	Perimeter of the cirlce is: %f
	`, c.area(), c.perimeter())
}
func printRectDetails(r rectangle) {
	fmt.Printf(`Area of Rectangle is: %f 
	Perimeter of the Rectangle is:  %f
	`, r.area(), r.perimeter())
}
func main() {
	perfectCircle := circle{5}

	myRect := rectangle{width: 5, height: 5}
	printCircleDetails(perfectCircle)
	printRectDetails(myRect)
}
