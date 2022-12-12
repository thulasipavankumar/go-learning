package main

import (
	"fmt"
)

var p, pf = fmt.Println, fmt.Printf

type Human struct {
	name string
	age  int
}

func changeValues(pointer *Human) {
	(*pointer).name = "uzumaki"
	pointer.age = pointer.age + 1
}
func main() {

	num := 42
	var pointer *int = &num
	fmt.Printf("%v,%v\n", pointer, *pointer)
	*pointer = 4242
	fmt.Println(num)
	fmt.Println(num == *pointer)
	var baby Human
	baby = Human{"baby", 1}
	fmt.Printf(" before modifying %#v\n", baby)
	changeValues(&baby)
	fmt.Printf(" before modifying %#v\n", baby)

}
