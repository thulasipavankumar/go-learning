package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("new.txt")
	if err != nil {
		log.Fatal("Unable to open file", err)
	}
	defer file.Close()
	byteSlice := make([]byte, 2)
	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	fmt.Println(numberBytesRead)
	fmt.Printf("%s\n", byteSlice)
	newfile, err := os.Open("new copy.txt")
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	data, err := ioutil.ReadAll(newfile)
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	fmt.Printf("%s : %T\n", data, data)
	fmt.Println(strings.Repeat("---", 10))
	read_by_line()
}
func read_by_line() {
	file, err := os.Open("new.txt")
	if err != nil {
		log.Fatal("Unable to open file", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	success := scanner.Scan()
	if success == false {
		err := scanner.Err()
		if err == nil {
			log.Println("Scan reached EOF")
		} else {
			log.Fatal("Error in reading the scanner", err)
		}
	}
	log.Println("first line is ", scanner.Text())

	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
func from_course() {
	//** READING INTO A BYTE SLICE USING io.ReadFull() **//

	// Opening the file in read-only mode. The file must exist (in the current working directory)
	// Use a valid path!
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// declaring a byte slice and initializing it with a length of 2
	byteSlice := make([]byte, 2)

	// io.ReadFull() returns an error if the file is smaller than the byte slice.
	// it reads the file into the byte slice up to its length
	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	fmt.Println(strings.Repeat("#", 20))

	//** READING WHOLE FILE INTO A BYTESLICE USING ioutil.ReadAll() **//

	// Opening another file (from the current working directory)
	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}

	// ioutil.ReadAll() reads every byte from the file and return a slice of unknown size
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))

	//** READING WHOLE FILE INTO MEMORY USING ioutil.ReadFile() **//

	// ioutil.ReadFile() reads a file into byte slice
	// this function handles opening and closing the file.
	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
}
