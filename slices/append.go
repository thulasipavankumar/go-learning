package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)
	fmt.Println(numbers[0:2]) // [1,2,3,4]
	fmt.Println(numbers[2:])
	fmt.Println(numbers)

}
