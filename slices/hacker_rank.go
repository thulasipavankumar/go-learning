package main

/*
https://www.hackerrank.com/challenges/plus-minus/problem?isFullScreen=true
*/
import (
	"fmt"
	"sort"
)

func main() {
	keyboard := []int32{4, 6, 5, 3, 3, 1}
	//	cpu := []int32{1, 2, 3, 5, 1}

	//plusMinus(arr)
	fmt.Println(pickingNumbers(keyboard))
}
func pickingNumbers(a []int32) int32 {
	// Write your code here
	sort.Sort(byValue(a))
	var longest int32 = 0
	for i := 0; i < len(a)-1; i++ {
		var streak int32 = 0
		var stream []int32
		stream = append(stream, a[i])
		for j := i + 1; j < len(a); j++ {
			if a[i]-a[j] == 1 || a[i]-a[j] == 0 || a[i]-a[j] == -1 {
				stream = append(stream, a[j])
				streak++
			} else {
				if longest < streak {
					fmt.Println(stream)
					longest = streak
				}

				break
			}
		}
	}
	if longest != 0 {
		return longest
	}
	return int32(-1)
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

type byValue []int32

func (f byValue) Len() int {
	return len(f)
}

func (f byValue) Less(i, j int) bool {
	return f[i] > f[j]
}

func (f byValue) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	var previous int32
	for _, i := range keyboards {
		for _, j := range drives {
			if i+j <= b && i+j > previous {
				previous = i + j
			}
		}
	}
	if previous != 0 {
		return previous
	}

	return int32(-1)
}
