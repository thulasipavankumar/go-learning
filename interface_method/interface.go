package main

import "fmt"

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return (c.radius * 3.14 * c.radius)
}
func main() {
	perfectCircle := circle{5}
	fmt.Println(perfectCircle.area())
	myRect := rectangle{width: 5, height: 5}
	_ = myRect
}
