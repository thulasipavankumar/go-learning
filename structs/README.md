## Structs 
```go
type book struct {
		title  string
		author string
		year   int
	}
	current_book := book{title:"The Subtle Art of Not Giving a F*ck", author:"Mark Manson", year: 2016}
    fmt.Printf("%+v\n", current_book)
```

## Compare structs
```go
type book struct {
		title  string
		author string
		year   int
	}
	current_book := book{title:"The Subtle Art of Not Giving a F*ck", author:"Mark Manson", year: 2016}
    last_book := book{title:"The Girl on the Train", author:"paula hawkins", year: 2015}
    fmt.Printf("%+v\n", current_book)
    fmt.Println(last_book == current_book)

```

## annonymus structs
```go
diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}
	fmt.Println(diana)
```

## annonymus struct fields
```go
    type Employee struct{
        string
        float64
        bool
    }

    myself := Employee{"me", 38000, true}
	fmt.Printf("%#v\n", myself)


     // mixing anonymous with named fields:
    type Employee1 struct {
        name   string
        salary int
        bool
    }

```

## Embedded structs
```go
    /** EMBEDDED STRUCTS **//
    // An embedded struct is just a struct that acts like a field inside another struct.
 
    // define a new struct type
    type Contact struct {
        email, address string
        phone          int
    }
 
    // define a struct type that contains another struct as a field
    type Employee struct {
        name        string
        salary      int
        contactInfo Contact
    }
    
    john := Employee{
        name:   "John Keller",
        salary: 3000,
        contactInfo: Contact{
            email:   "jkeller@company.com",
            address: "Street 20, London",
            phone:   042324234,
        },
    }
```