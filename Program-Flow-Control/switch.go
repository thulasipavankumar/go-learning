package main

import "fmt"

func main() {

	var language string = "Golang"
	switch language {
	case "Python":
		fmt.Println("you are learning pythin")
	case "Golang":
		fmt.Println("you are learning golang")
	case "C++":
		fmt.Println("you are learning C++")
	case "Java":
		fmt.Println("you are learning java")
	}
	condition := 1
	switch condition {
	case 1, 2:
		fmt.Println("case 1 & case 2") // two cases have same block of code for execution
	case 3:
		fmt.Println("case 3")
	default:
		fmt.Println("default case")
	}

}
