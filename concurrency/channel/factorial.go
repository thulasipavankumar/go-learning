package main

import (
	"fmt"
	"math/big"
)

func factorial(n int, ch chan *big.Int) {
	fact := big.NewInt(1)

	for i := 2; i <= n; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}
	ch <- fact
	fmt.Println("Completed Factorial for: ", n)
}

func main() {
	//slice :=[]int{10,9,8,7,6,5,4,3,2,1}
	const limit int = 10
	ch := make(chan *big.Int, limit) // ==> unbuffered channel
	//ch := make(chan *big.Int) //=> buffered channel
	defer close(ch)
	for i := range [limit]int{} {
		go factorial(i, ch)
	}
	for i := range [limit]int{} {
		fmt.Println(i, <-ch)
	}

}
