package main

import "fmt"

//import "fmt"

func main() {
	var x = 10
	fmt.Println(x)

	//  var y = 15;  this will return compile error because y is not used , but can be overcome using underscore
	var unused = "blank"
	_ = unused

	name, age := "Pavan Kumar Tulasi", 29
	fmt.Println(name, age)
	fmt.Println("year passed")

	age, gender := 30, "Male"
	_ = gender
	fmt.Println(name, age)

	var i, j, k int
	fmt.Println(i, j, k)
	// mutiple declaration for readability
	var (
		pet     string
		pet_age int
		cost    float64
	)

	_, _, _ = pet, pet_age, cost

	//swap assignemt
	i, j = 1, 2
	i, j = j, i

	fmt.Println(i, j)

}
