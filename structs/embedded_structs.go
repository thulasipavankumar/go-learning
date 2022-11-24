package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}
	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}
	naruto := Employee{
		name:   "Naruto",
		salary: 4000,
		contactInfo: Contact{
			email:   "Naruto@konoha.com",
			address: "leaf Village",
			phone:   7896,
		},
	}
	fmt.Printf("%#v\n", naruto)
	fmt.Printf("Employee name %s and contact is %s", naruto.name, naruto.contactInfo.email)
}
