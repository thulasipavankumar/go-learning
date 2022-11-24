package main

import "fmt"

func main() {
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}
	fmt.Println(diana)
	type Employee struct {
		string
		float64
		bool
	}
	var myself Employee
	myself = Employee{"me", 38000, true}
	fmt.Printf("%#v\n", myself)
}
