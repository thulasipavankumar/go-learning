package main

import "fmt"

type names []string

func main() {
	characters := names{"Naruto", "Sakura", "Sasuke"}
	characters.print()
}
func (n names) print() {
	for _, name := range n {
		fmt.Println(name)
	}
}
