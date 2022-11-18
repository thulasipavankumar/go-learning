package main

/*
https://www.hackerrank.com/challenges/plus-minus/problem?isFullScreen=true
*/
import "fmt"

func main() {
	arr := []int32{1, 2, 3, -2, -1}
	plusMinus(arr)
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
