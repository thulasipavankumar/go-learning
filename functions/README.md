## Basic Functions
```go
func add(a, b int) (sum int) {
    sum = a+b
	return   // Naked return , it will return sum because it's already defined sum is return type in func decalartion
}

// defining a function that have two parameters of type int and returns two values of type int

func add_mutiply(a, b int) (int, int) {
    return a * b, a + b
}

mul, sum := add_mutiply(42, 42)


// defining a function that have one parameter of type float64 and returns a value of type float64
func pow(a float64) float64 {
    return math.Pow(a, a)
    //any statements below the return statement are never executed
 
}
```

## There are no default arguments in go

## variadic functions
```go
func variadic(nums ...int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}

```

### anonymus functions
```go
func main() {
	func(message string) {
		fmt.Println("Hello! ", message)
	}("Naruto")
}

func increment(x int) func() int {
    return func() int {
        x++
        return x
    }
}

    a := increment(10)
    fmt.Println(a()) // 11
    fmt.Println(a()) // 12
    fmt.Println(a()) // 13
```


### defer functions
```go
package main
 
import "fmt"
 
func f1() {
    fmt.Println("Calling f1")
}
 
func f2() {
    fmt.Println("Calling f2")
}
 
func main() {
    defer f1()
    f2()
}
/* OUTPUT
Calling f2
Calling f1
*/
```