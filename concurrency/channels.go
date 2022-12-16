package main

import "fmt"

func write_to_channel(n int, ch chan int) {
	ch <- n
}
func main() {
	//var ch chan int
	ch := make(chan int, 1)
	go write_to_channel(42, ch)
	n := <-ch
	fmt.Println(n)
}
