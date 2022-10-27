# go-learning


## To run  go program

 "go run < gofile >" 

example:- `go run main.go`


## To print value

use println func under fmt package

example:- 
``` 
fmt.Println("Hello World!")
name := "naruto"
fmt.Println(name)
```
## declared-and-not-used error referenece
 If a variable is decalted but not used anywhere go will throw error.
 To avoid this use _ as left assignment to the unsed value
example:- 
``` 
fmt.Println("Hello World!")
name := "naruto"
i=1
_=i
fmt.Println(name)
```
