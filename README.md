# go-learning


## To run  go program

 "go run < gofile >" 

example:- `go run main.go`


## To print value

use println func under fmt package

example:- 
```go 
fmt.Println("Hello World!")
name := "naruto"
fmt.Println(name)
```
## declared-and-not-used error referenece
 If a variable is decalted but not used anywhere go will throw error.
 To avoid this use _ as left assignment to the unsed value
example:- 
```go 
fmt.Println("Hello World!")
name := "naruto"
i=1
_=i
fmt.Println(name)
```

## mutiple declaration for readability
```go
	var (
		pet     string
		pet_age int
		cost    float64
	)
```

## Constants

 You cannot initiate a constant at runtime (constants belong to compile-time)
 const power = math.Pow(2, 3) //error, functions calls belong to runtime

Initializing the constants using a step:
 ```go
     const (
        c11 = iota * 2 // -> 0
        c22            // -> 2
        c33            // -> 4
    )
 ```

```go
  var i3 int64 = -324_567_345  // underscores are used to write large num
```

Unlike c,c++,java when increment or drecement is done it is stored

```go
x:=1
x++
fmt.Println(x) // 2

var i int8 = 127;

i++ // overflows and no error

j:= int8(255+1)  // throws overflow error
```

## Converting Numbers (or) Strings
```go

x:-1
y:=2.1
x = x+int(y) // y is converted to int for addition but y is still float

// Numbers to String
s = string(99)   // =>  s=c because 97 is ascii for character 'c'

s = fmt.Sprintf("%f",3.14) // use fmt.Sprintf for converting float to string

s = fmt.Sprintf("%d",534)  // use fmt.Sprintf for converting int  to string  => s = 534 
// !! above there string is not the ascii value but the number itself

// String to Numbers
const pi_string = "3.14"

 pi_string,err := strconv.ParseFloat(pi,64)   // parse float64

i,err:= strconv.Itoa(20)  // => i = "20"  number to string

i,err := strconv.Atoi("-50")  // ==> i = -50  string to number

```


## Defined type and alias

```go
type speed uint
 
// s1, s2 of type speed
var s1 speed = 10   // type of s1 is speed and undelying type is uint
var s2 speed = 20
 
// performing operations with the new types
fmt.Println(s2 - s1) // -> 10
 
// uint and speed are different types (they have different names)
var x uint
 
// x = s1  //error different types


// Alias
type second = uint
 
var hour second = 3600   //no need to convert operations (same type)
fmt.Printf("hour type: %T\n", hour) // => hour type: uint
 
    
```

## Command line arguments

```go
fmt.Printf("Arguments are %v and length of arguments are %d", os.Args,len(os.Args))
//zero index is the go executable path
```
## There is no while loop in go