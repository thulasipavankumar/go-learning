package main

import (
	"fmt"
	"os"
)

func main() {
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)
	fmt.Println(numbers[0:2]) // [1,2,3,4]
	fmt.Println(numbers[2:])
	fmt.Println(numbers)
	args()

}
func args() {

	arguments := os.Args
	arguments[0] = "myrandom.exe"
	fmt.Println(os.Args)
	_ = arguments
}
