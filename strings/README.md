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
```
