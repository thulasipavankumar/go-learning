package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)
	fmt.Println(numbers[0:2]) // [1,2,3,4]
	fmt.Println(numbers[2:])
	fmt.Println(numbers)
	args()
	assignment()

}
func args() {

	arguments := os.Args
	arguments[0] = "myrandom.exe"
	fmt.Println(os.Args)
	_ = arguments
}
func assignment() {
	arguments := os.Args[1:]
	var add, multiply int
	multiply = 1
	for _, v := range arguments {
		temp, _ := strconv.Atoi(v)
		add += temp
		multiply *= temp
	}
	fmt.Println(add, multiply)
}
