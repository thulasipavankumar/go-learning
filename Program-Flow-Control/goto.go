package main

import "fmt"

func main() {
	count := 0

loop:
	if count < 10 {
		fmt.Println(count)
		count++
		goto loop
	}
	fmt.Println("Exited the loop")

	fmt.Println("Reached Program end")

}
