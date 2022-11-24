package main

import (
	"log"
	"os"
)

const file_name = "info.txt"

var p, pf, f = log.Println, log.Printf, log.Fatal

func main() {
	assign1()
}

func assign1() {
	/*
	   Create a new file in the current working directory called info.txt and then close the file.
	   If the file cannot be created (no permissions, wrong path etc)
	   then print out the error and stop the program (do error handling).
	*/

	file, err := os.Create(file_name)
	if err != nil {
		f("Fatal error in creating the file", err)
	}
	defer file.Close()
	_ = file
}

func assign2() {
	/*
		Rename the file created at Exercise #1 to data.txt

		Check if file exists before renaming it. If it doesn't exist print a message and stop the program.
	*/
	_, err := os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			f("The file does not exist!")
		}

	}
	err = os.Rename("info.txt", "data.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
func assign3() {
	err := os.Remove("data.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
func assign4() {
	/*
		Create a Go Program that reads the entire contents of a file called info.txt into a string.
		You can use ioutil.ReadAll() or any other function you want.

		The file exists in the current working directory.
	*/
	file, err := os.Open(file_name)
	//buff, err := ioutil.ReadAll()
	_, _ = file, err

}
