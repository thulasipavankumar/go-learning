## Declaring pointer
```go
    num := 42
	var pointer *int = &num
    fmt.Printf("%v,%v",pointer,*pointer) // => prints the address and value 0xc000106018,42
    *pointer = 4242
	fmt.Println(num)  // => 4242 because the value is changed by the pointer reference
    
```

```go
    var x int = 2
    // the expression &x means the address of x and creates a pointer to an integer variable,
    // ptr is of type *int, which is pronounced "pointer to int".
    ptr := &x
    fmt.Printf("ptr is of type %T with value %v and address %p\n", ptr, ptr, &ptr) // -> p is of type *int with value 0xc000014140.
    
    p := new(int) // it creates a pointer called p that is a pointer to an int type
 
    x = 100
    p = &x // initializing p
    *p = 90 //equivalent to x = 90 because *p is x
    
    // x and *p is the same thing.
 
    fmt.Println(*p)                  // => 90
    fmt.Println("*p == x:", *p == x) // => *p == x: true
```

## pointers on struct
```go

type Human struct {
	name string
	age  int
}

func changeValues(pointer *Human) {
	(*pointer).name = "uzumaki"
	pointer.age = pointer.age + 1
}
func main() {

	
	var baby Human
	baby = Human{"baby", 1}
	fmt.Printf(" before modifying %#v\n", baby) //  before modifying main.Human{name:"baby", age:1}
	changeValues(&baby)
	fmt.Printf(" before modifying %#v\n", baby) //  before modifying main.Human{name:"uzumaki", age:2}

}
```

## Slices and map are always pass by pointer
```go
func changeSlice(s []int) {
	for i := range s {
		s[i] += 1
	}
}
func changeMap(m map[string]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}
func main(){
    slice := []int{3, 4, 5}

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Printf("before change slice %#v\n", slice) //   before change slice []int{3, 4, 5}
	fmt.Printf("after change map %#v\n", m) //   after change map map[string]int{"a":1, "b":2, "c":3}
	changeMap(m)
	changeSlice(slice)
	fmt.Printf("before change slice %#v\n", slice) //   before change slice []int{4, 5, 6}
	fmt.Printf("after change map %#v\n", m)  //     after change map map[string]int{"a":2, "b":3, "c":4}
}
```
