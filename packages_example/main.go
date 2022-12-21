package main

import "fmt"
// Below package is located in "/usr/local/go/src/mypackages/arithmetic"
import "mypackages/arithmetic"
func init() {
	fmt.Println("This is init block")
}
func main() {
	num:=13
	fmt.Println("This is main block")
	fmt.Printf("Is the number %d prime: %v\n",num,arithmetic.IsPrime(num))
}
