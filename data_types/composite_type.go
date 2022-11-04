package main

import "fmt"

func main() {

	// ARRAY Type
	var days = [7]string{"Mon", "Tues", "Wed", "Thrus", "Fri", "Sat", "Sun"}
	fmt.Printf("%T\n", days)

	// SLICE Type    => compared to Array SLICE has dynamic size

	var num = []int{1, 2, 3, 4, 5}

	//Map

	var balance = map[string]float64{
		"USD": 121.54,
		"EUR": 10000,
	}
	fmt.Println(balance)

	// Struct
	type Person struct {
		name string
		age  int
	}
	var me Person

	// Pointer type

	var x int = 5
	ptr := &x
	fmt.Println(x, ptr)

	//Function type
	fmt.Printf("%T\n", f)
	_, _ = num, me
}
func f() {

}
