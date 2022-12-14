package main

import "fmt"

type emptyInterface interface {
}

func main() {
	var empty emptyInterface
	empty = 10
	fmt.Println(empty)
	empty = "hello world"
	fmt.Println(empty)
	empty = []int{1, 2, 3}
	fmt.Println(len(empty.([]int)))

}
