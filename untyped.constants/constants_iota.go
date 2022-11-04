package main

import "fmt"

func main() {
	const (
		c1 = iota // 0
		c2 = iota // 1
		c3 = iota // 2
		c4 = iota // 3
	)
	const name = "Pavan Kumar Tulasi"

	const (
		d1 = iota // 0
		d2        // 1
		d3        // 2
		d4        // 3
	)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println(d1, d2, d3, d4)
	const (
		zero  = iota // 0
		_            // 1
		two          // 2
		three        // 3
	)
	fmt.Println(zero, two, three)
}
