## Methods(Reciever functions) can be defined on types
## Methods can only be created for value types not pointer types
```go
  type names* []string 
  func (n names) print() {}  // ==> throw error because method is creted for pointer type
```

- function belongs to the package
- method belongs to the type

```go
type names []string


func main() {
	characters := names{"Naruto", "Sakura", "Sasuke"}
	characters.print()  // => this will print out the names as defined in the below method
}

func (n names) print() {
	for _, name := range n {
		fmt.Println(name)
	}
}
```