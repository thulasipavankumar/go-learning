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
