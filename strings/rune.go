package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(len("€"), utf8.RuneCountInString("€")) // len = 3 , rune =1
	s2 := "中文维基是世界上"
	fmt.Println(s2[0:2])
	rs := []rune(s2)
	fmt.Println("string slice is ", s2[0:2])
	for _, v := range rs {
		fmt.Print(string(v))
	}
}
