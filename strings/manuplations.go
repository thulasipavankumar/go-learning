package main

import (
	"fmt"
	"strings"
)

func main() {
	println := fmt.Println
	name := "Naruto"
	println(strings.Contains(name, "Nar"))
	println("Hello")
}
