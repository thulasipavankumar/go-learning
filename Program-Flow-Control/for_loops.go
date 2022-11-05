package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	i := 10
	//  similar to while loop but using for
	for i >= 0 {
		fmt.Println(i)
		i--
	}
}
