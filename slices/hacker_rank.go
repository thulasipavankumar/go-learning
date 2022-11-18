package main

/*
https://www.hackerrank.com/challenges/plus-minus/problem?isFullScreen=true
*/
import "fmt"

func main() {
	//arr := []int32{1, 2, 3, -2, -1}

	//plusMinus(arr)
	fmt.Println(catAndMouse(3, 4, 5))
}

func plusMinus(arr []int32) {
	// Write your code here
	postivie, negative, zero := 0.0, 0.0, 0.0
	for _, v := range arr {
		if v > 0 {
			postivie++
		} else if v < 0 {
			negative++
		} else {
			zero++
		}

	}
	postive_ratio, negative_ratio, zero_ratio := postivie/float64(len(arr)), negative/float64(len(arr)), zero/float64(len(arr))
	fmt.Printf("%.6f\n%.6f\n%.6f\n", postive_ratio, negative_ratio, zero_ratio)

}

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	a, b := abs(z-x), abs(z-y)
	if a == b {
		return "Mouse C"
	}
	if a < b {
		return "Cat A"
	}
	return "Cat B"

}
func abs(c int32) int32 {
	if c < 0 {
		return c * -1
	}
	return c
}
