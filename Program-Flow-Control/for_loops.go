package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i <= 10; i++ {
		//fmt.Println(i)
	}
	i := 10
	//  similar to while loop but using for
	for i >= 0 {
		//fmt.Println(i)
		i--
	}
	i = 11
	//continue & break for printing even number
	for {
		i--
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
		if i == 0 {
			break
		}
	}

	for index, value := range os.Args {

		fmt.Printf("Value %q is a index %d and types are %T ,%T\n", value, index, value, index)
	}

}
