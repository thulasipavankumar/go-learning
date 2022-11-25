package main

import "fmt"

func main() {
	func(message string) {
		fmt.Println("Hello! ", message)
	}("Naruto")
	fmt.Println(increment(10)(11))
}

func increment(x int) func(val int) int {
	return func(val int) int {
		return x + val
	}
}
