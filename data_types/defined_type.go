package main

import "fmt"

type km float64
type mile float64

func main() {

	type age int
	type adult age  // int is underlying type
	type infant age //  int is underlying type

	type speed uint
	var s1 speed = 10
	var s2 speed = 20
	fmt.Println(s2 - s1)

	var x uint

	//x= s1  // compile time error have to convert

	x = uint(s1)

	var s3 speed = speed(x)

	_, _ = x, s3

	var paris_to_brussels km = 465
	var distance_in_miles mile
	//distance_in_miles = paris_to_brussels / 0.621 // => error because km is assigned to mile type
	distance_in_miles = mile(paris_to_brussels)
	_ = distance_in_miles

}
