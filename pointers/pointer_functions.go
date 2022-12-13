package main

import "fmt"

func change(x *int) *float64 {
	*x = 10000
	y := 4.2
	return &y
}
func main() {
	i := 1
	p := &i
	fmt.Println(change(p))
	fmt.Println(i)
	slice := []int{3, 4, 5}

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Printf("before change slice %#v\n", slice)
	fmt.Printf("after change map %#v\n", m)
	changeMap(m)
	changeSlice(slice)
	fmt.Printf("before change slice %#v\n", slice)
	fmt.Printf("after change map %#v\n", m)
}
func changeSlice(s []int) {
	for _, i := range s {
		s[i] += 1
	}
}
func changeMap(m map[string]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}
