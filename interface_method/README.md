## Methods(Reciever functions) can be defined on types
## Methods can only be created for value types not pointer types
## A type assertion provides access to an interface’s dynamic concrete value.
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



```go
import (
    "fmt"
    "math"
)
 
// declaring an interface type called shape
// an interface contains only the signatures of the methods, but not their implementation
type shape interface {
    area() float64
    perimeter() float64
}
 
// declaring 2 struct types that represent geometrical shapes: rectangle and circle
 
type rectangle struct {
    width, height float64
}
type circle struct {
    radius float64
}
 
// method that calculates circle's area
func (c circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}
 
// method that calculates rectangle's area
func (r rectangle) area() float64 {
    return r.height * r.width
}
 
// method that calculates circle's perimeter
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
 
// method that calculates rectangle's perimeter
func (r rectangle) perimeter() float64 {
    return 2 * (r.height + r.width)
}
 
// any type that implements the interface is also of type of the interface
// rectangle and circle values are also of type shape
func print(s shape) {
    fmt.Printf("Shape: %#v\n", s)
    fmt.Printf("Area: %v\n", s.area())
    fmt.Printf("Perimeter: %v\n", s.perimeter())
}
 
func main() {
 
    // declaring a circle and a rectangle values
    c1 := circle{radius: 5}
    r1 := rectangle{width: 3, height: 2}
 
    // circle and rectangle both implements the geometry interface  because they implement all methods of the interface
    // an interface is implicitly implemented in Go
    print(c1)
    print(r1)
 
    a := c1.area()
    fmt.Println("Circle Area:", a)
}
```

# Type Assertions, Type Switch google !!!

```go
package main
 
import (
    "fmt"
    "math"
)
 
// declaring an interface type called shape
type shape interface {
    area() float64
    perimeter() float64
}
 
// declaring a struct type
type rectangle struct {
    width, height float64
}
 
// declaring a struct type
type circle struct {
    radius float64
}
 
// declaring a method for circle type
func (c circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}
 
// declaring a method for circle type
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
 
// declaring a method for circle type
func (c circle) volume() float64 {
    return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}
 
// declaring a method for rectangle type
func (r rectangle) area() float64 {
    return r.height * r.width
}
 
// declaring a method for rectangle type
func (r rectangle) perimeter() float64 {
    return 2 * (r.height + r.width)
}
 
// declaring a function that takes an interface value
func print(s shape) {
 
    fmt.Printf("Shape: %#v\n", s)
    fmt.Printf("Area: %v\n", s.area())
    fmt.Printf("Perimeter: %v\n", s.perimeter())
}
 
func main() {
 
    // declaring an interface value that holds a circle type value
    var s shape = circle{radius: 2.5}
 
    fmt.Printf("%T\n", s) //interface dynamic type is circle
 
    // no direct access to interface's dynamic values
    // s.volume() -> error
 
    // there is access only to the methods that are defined inside the interface
    fmt.Printf("Circle Area:%v\n", s.area())
 
    // an interface value hides its dynamic value.
    // use type assertion to extract and return the dynamic value of the interface value.
    fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())
 
    // checking if the assertion succeded or not
    ball, ok := s.(circle)
    if ok == true {
        fmt.Printf("Ball Volume:%v\n", ball.volume())
    }
 
    //** TYPE SWITCHES **//
 
    // it permits several type assertions in series
    switch value := s.(type) {
    case circle:
        fmt.Printf("%#v has circle type\n", value)
    case rectangle:
        fmt.Printf("%#v has rectangle type\n", value)
 
    }
}
```

## embedded interface contains other interfaces in it
```go
```

## empty interface can hold any value
```go
type emptyInterface interface {
}

func main() {
	var empty emptyInterface
	empty = 10
	fmt.Println(empty)  // => 10 
	empty = "hello world"
	fmt.Println(empty)  // => hello world 
	empty = []int{1, 2, 3}
	fmt.Println(len(empty.([]int)))  // => type assertions to use len method

}
```