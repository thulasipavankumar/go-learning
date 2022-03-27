package main

import (
	"fmt"
)

// This is a comment
func main() {
	var a int = 5
	var b = 3.4           // type inference
	name := "Pavan kumar" // short declaration
	_ = name              // Blank Identifier (_) to get rid of unused variable
	fmt.Println(a, b)
	a = int(b)
	fmt.Println(a, b)
}

/*
  this is comment block
    ^  _  ^
*/
