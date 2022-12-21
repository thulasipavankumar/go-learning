package main

import "fmt"
// Below package is located in "/usr/local/go/src/mypackages/arithmetic"
import local_arithmetic "mypackages/arithmetic"
import git_arithmetic "github.com/thulasipavankumar/go-mod-math/arithmetic"
//import v1 "github.com/thulasipavankumar/go-mod-math/arithmetic@v1.1.0"
func init() {
	fmt.Println("This is init block")
}
func main() {
	num:=13
	fmt.Println("This is main block")
	fmt.Printf("Is the number %d prime: %v\n",num,local_arithmetic.IsPrime(num))
	num = 10
	fmt.Printf("Is the number %d prime: %v\n",num,git_arithmetic.IsPrime(num))
	fmt.Println(git_arithmetic.Factorial(num))
	//fmt.Printf("Factorial of %d is: %v\n",num,v1.Factorial(num))

	
}
