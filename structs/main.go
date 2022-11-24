package main

import (
	"fmt"
	"log"
)

var p, pf, f = log.Println, log.Printf, log.Fatal

type book struct {
	title  string
	author string
	year   int
}
type person struct {
	firstName, lastName string
	age                 int
}

func main() {
	title, author, year := "The Subtle Art of Not Giving a F*ck", "Mark Manson", 2016
	pf("%s %s %d \n", title, author, year)

	current_book := book{title: "The Subtle Art of Not Giving a F*ck", author: "Mark Manson", year: 2016}
	p(current_book)
	lastBook := book{title: "iam title", author: "iam the author", year: 1}
	p(lastBook.title)
	lastBook.year = 2
	p(lastBook.year)
	fmt.Printf("%+v\n", current_book)
	last_book := book{title: "The Girl on the Train", author: "paula hawkins", year: 2015}
	fmt.Printf("%+v\n", current_book)
	fmt.Println(last_book == current_book)
}
func course() {

	// creating a struct type
	type book struct {
		title  string //the fields of the book struct
		author string //each field must be unique inside a struct
		year   int
	}

	// combining different fields of the same type on the same line
	type book1 struct {
		title, author string
		year, pages   int
	}

	// declaring, initializing and assigning a new book value, all in one step
	lastBook := book{"The Divine Comedy", "Dante Aligheri", 1320} //this is a struct literal and order matters
	fmt.Println(lastBook)

	// Declaring a new book value by specifying field: value (order doesn't matter)
	bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945}
	_ = bestBook

	//if we create a new struct value by omitting some fields they will be zero-valued according to their type
	aBook := book{title: "Just a random book"}
	fmt.Printf("%#v\n", aBook) // => main.book{title:"Just a random book", author:"", year:0}

	// retrieving the value of a struct field
	fmt.Println(lastBook.title) // => The Divine Comedy

	// selecting a field that doesn't exist raises an error
	// pages := lastBook.pages // error -> lastBook.pages undefined (type book has no field or method pages)

	// updating a field
	lastBook.author = "The Best"
	lastBook.year = 2020
	fmt.Printf("lastBook: %+v\n", lastBook) // => lastBook: {title:The Divine Comedy author:The Best year:2020}
	// + modifier with %v  printed out both the field names and their values

	// comparing struct values
	// two struct values are equal if their corresponding fields are equal.
	randomBook := book{title: "Random Title", author: "John Doe", year: 100}
	fmt.Println(randomBook == lastBook) // => false

	// = creates a copy of a struct
	myBook := randomBook
	myBook.year = 2020              // modifying only myBook
	fmt.Println(myBook, randomBook) // => {Random Title John Doe 2020} {Random Title John Doe 100}
}
