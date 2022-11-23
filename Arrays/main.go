package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	even_numbers := [5]int{2, 4, 6, 8, 10}
	_ = even_numbers
	currency := [...]string{"INR", "USD", "EUR"}
	fmt.Println(currency)
	full_name := [...]string{
		"Pavan",
		"kumar",
		"Tulasi", // for multiple line decalarion ',' should be present at the last element
	}
	_ = full_name
	fmt.Println(strings.Repeat("#", 10))
	for i, v := range currency {
		fmt.Printf("at index %d value is %v \n", i, v)
	}
	fmt.Println(strings.Repeat("#", 10))
	for i := 0; i < len(even_numbers); i++ {
		fmt.Println(even_numbers[i])
	}

	muti_dimension_array := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	_ = muti_dimension_array
	keyed_elements()
}
func keyed_elements() {
	random := [3]int{
		1: 67,
		0: 8,
		2: 17,
	}
	fmt.Println(random)
	fmt.Println(strings.Repeat("#", 10))
	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend)
	fmt.Println("My favorite number is", rand.Intn(100))
}
