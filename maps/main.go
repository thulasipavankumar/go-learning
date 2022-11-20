package main

import (
	"fmt"
	"strings"
)

func main() {
	pf := fmt.Printf
	p := fmt.Println
	var employess map[string]string
	p("No.of pairs:%d", len(employess))
	p("Employe value of Naruto is ", employess["Naruto"])
	pf("%#v\n", employess)
	employess = map[string]string{}
	employess["Naruto"] = "Hokage"
	p("Employe value of Naruto is ", employess["Naruto"])

	balances := map[string]float64{
		"USD": 111.5,
		"EUR": 10000,
		"GBP": 15,
	}
	fmt.Println(balances)
	balances["INR"] = 10000
	v, ok := balances["CHF"]
	if ok {
		fmt.Println("CHF balance exsists", v)
	} else {
		fmt.Println("Given balance doesnot exsist")
	}

	for k, v := range balances {
		p(k, v)
	}
	delete(balances, "GBP")
	_ = balances
	p(strings.Repeat("#", 10))
	comparisions()
}

func comparisions() {
	a := map[string]string{"A": "0", "B": "1"}
	b := map[string]string{"Y": "24", "Z": "25"}
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are  not equal")
	}
}
