package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(200)
	n := 0
	for i := 100; i > 0; i-- {
		go func() {
			m.Lock()
			n++
			m.Unlock()
			//fmt.Println("Increament:", n)
			wg.Done()
		}()
		go func() {
			m.Lock()
			n--
			m.Unlock()
			//fmt.Println("Decrement:", n)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n)
	_ = m

}
