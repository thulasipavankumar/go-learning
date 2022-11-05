package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// IF Statement

	//name, age := "Naruto", 19

	// Command line Arguments

	fmt.Println("Arguments are", os.Args)
	//fmt.Println(strconv.ParseFloat(os.Args[1], 64))
	//var name string
	if args := os.Args; len(args) < 3 {

		fmt.Println("Please provide 2 inputs => first name and age(float)")
	} else if age, err := strconv.ParseFloat(args[2], 64); err == nil {
		name := args[1]
		fmt.Printf("User %q is of age %f\n", name, age)
	} else {
		fmt.Println(" please provide number as age for the 2nd argument", err)
	}
}
