package main

import "fmt"

func main() {
	people := [5]string{"Hinata", "Sakura", "Kakashi", "Naruto", "Sasuke"}
	friends := [2]string{"Naruto", "Sasuke"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d", name, index)
				break outer
			}
		}

	}
	fmt.Println("Entered here")

}
