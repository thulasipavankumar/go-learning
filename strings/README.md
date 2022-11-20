## String in go

Strings are sequence of bytes not characters in GO

### There is no charachers in go , only byte(uint8) and rune(inte32) data types
```go
string_name := "Naruto"
string_alias := `Hokage`
execution_path := `c:\users\documents\main.go`
execution_path_escape_char := "c:\\users\\documents\\main.go"

```

```go
fmt.Println(len("€"), utf8.RuneCountInString("€"))  // len = 3 , rune =1 
s2 := "中文维基是世界上"
    fmt.Println(s2[0:2]) // -> � - the unicode representation of bytes from index 0 and 1.
 
    // returning a slice of runes
    // 1st step: converting string to rune slice
    rs := []rune(s2)
    fmt.Printf("%T\n", rs) // => []int32
    // 2st step: slicing the rune slice
    fmt.Println(string(rs[0:3])) // => 中文维
```

### String manuplations
```go

// it returns a copy of a string by replacing a substring (old) by another substring (new)
str := strings.Replace("192.168.0.1", ".", ":", 2)     // -> 192:168:0.1   it replaces the first 2 occurrences

str := strings.Replace("192.168.0.1", ".", ":", -1)     // -> 192:168:0:1

str := strings.ReplaceAll("192.168.0.1", ".", ":")      // -> 192:168:0:1

strings.Count("cheese", "e") // => 3

 // if the substr is an empty string Count() returns 1 + the number of runes in the string
 strings.Count("five", "")// => 5 (1 + 4 runes)

 str := strings.ToLower("Go Python Java") // -> go python java

 str := strings.ToUpper("Go Python Java") // -> GO PYTHON JAVA

 str := strings.split("192.168.0.1", ".")  // =>  [192,168,0,1]

 str := strings.Split("Go for Go!", "")   // => []string{"G", "o", " ", "f", "o", "r", " ", "G", "o", "!"}

  // Join() concatenates the elements of a slice of strings to create a single string.
  s = []string{"I", "learn", "Golang"}
 
  str := strings.Join(s, "-")   // -> I-learn-Golang

  // splitting a string by whitespaces and newlines.

  str := strings.Fields("Hello world!\n I'm Alive")  // => [Hello world! I'm Alive]

  // TrimSpace() removes leading and trailing whitespaces and tabs.

  str := strings.TrimSpace("\t Goodbye Windows, Welcome Linux!\n ") // =>  "Goodbye Windows, Welcome Linux!"

   // To remove other leading and trailing characters, use Trim()

  str := strings.Trim("...Hello, Gophers!!!?", ".!?")   // "Hello, Gophers"
 ```
### String comparisions

```go
    strings.Contains("hello", "hell")   // => true

    strings.ContainsAny("hello","how")  // => true hello contains h and also o

    strings.ContainsRune("golang", 'g') // => true

 // comparing strings (case matters)
    "go" == "go" // -> true
    "Go" == "go"  // -> false

    strings.ToLower("Go") == strings.ToLower("go") // -> true

 // EqualFold() compares strings (case doesn't matter) -> it's efficient
    strings.EqualFold("Go", "gO") // -> true


```