package main

import "fmt"

func main() {
	factions := []string{"leaf village", "sand village", "Konohagakure", "Uzushiogakure", "Amegakure"}
	favourites, others := factions[:1], factions[1:]

	fmt.Println(favourites, others)
	others[0] = "Sunagakure"
	fmt.Println(factions)
	sequence := [5]int{1, 2, 3, 4, 5}
	slice1, slice2 := sequence[:2], sequence[2:]
	slice1[0] = 10
	_ = slice2
	fmt.Println(sequence)
	fmt.Println(len(slice1), cap(slice1)) // len 2 , capacity 5
	fmt.Println(len(slice2), cap(slice2)) // len 3 , capacity 4 because the slice is created from arr[2] index
	part2()
}
func part2() {
	cars := []string{"maserati", "vw", "ferrari"}
	newCars := []string{}
	newCars = append(newCars, cars...) // since we used append they donot have same backing array any more
	// change 'newCars' will not affect 'cars' slice

}
