## IF ELSE caluse

```go
if age>=18 {  
    fmt.println("You are an Adult")
} else {
    fmt.Println("You are a Minor")
}
```

### Strange part is the brace has to be in the same line as IF otherwise go throws error


## Command line arguments

```go
fmt.Printf("Arguments are %v and length of arguments are %d", os.Args,len(os.Args))
//zero index is the go executable path
```

## There is no while loop in go

## break statements are automaticly assumed in switch case no need to explicitly mention them
```go

switch condition{
    case 1:
    fmt.Println("case 1")   // => no need to explicitly mention break 
    case 2:
    fmt.Println("case 2")
    case 3:
    fmt.Println("case 3")
    defulat:
    fmt.Println("default case")
}
switch condition{
    case 1, 2:
    fmt.Println("case 1 & case 2")  // two cases have same block of code for execution
    case 3:
    fmt.Println("case 3")
    default:
    fmt.Println("default case")
}
 switch n := 10;true{
    case n > 0:
        fmt.Println("Positive")
    case n < 0:
        fmt.Println("Negative")
    default:
        fmt.Println("Zero")
    }
```
