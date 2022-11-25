package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	example()
	greet("kuruma")
	p(add(42, 42))
	mul, sum := add_mutiply(42, 42)
	p(mul, sum)
	variadic(1, 2, 3, 4)
	info := personInformation(35, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info) // => Age: 35, Full Name:Wolfgang Amadeus Mozart
}
func example() {
	p("This is a function")
}
func greet(message string) {
	p("Hello ", message)
}
func add(a, b int) (sum int) {
	sum = a + b
	return sum
}
func add_mutiply(a, b int) (int, int) {
	return a * b, a + b
}

func variadic(a ...int) {
	for _, v := range a {
		p(v)
	}
}

// mixing variadic and non-variadic parameters is allowed
// non-variadic parameters are always before the variadic parameter
func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name:%s", age, fullName)
	return returnString
}
