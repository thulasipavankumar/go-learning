package main

import (
	"fmt"
	"strconv"
)

const name = "Naruto Uzumaki" // package scoped
func main() {
	const name = "dattebayo "
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(name) // => name = "dattebayo "  due to local scope reference
	sample()
}
func sample() {
	fmt.Println(name) // => name = "Naruto Uzumaki"  due to global scope reference
}
