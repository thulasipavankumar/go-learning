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
}

func main() {
	//slice :=[]int{10,9,8,7,6,5,4,3,2,1}
	const limit int = 200
	ch := make(chan *big.Int, limit)
	defer close(ch)
	for i, _ := range [limit]int{} {
		go factorial(i, ch)
	}
	for i, _ := range [limit]int{} {
		fmt.Println(i, <-ch)
	}

}
