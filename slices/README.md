## Bottom some examples of creating slices
```go
var cities []string // empty slice
var numbers = []int{1, 2, 3, 4, 5}
var slice = make([]string, 5) // slice declaration with size


    type number []int
	even := number{2, 4, 6, 8, 10}
	odd := number{1, 3, 5, 7, 9, 11}
```

## appending the slices
```go
    numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)
    negative := []int{-3,-2,-1}
    integers:=apeend(negative,numbers...)
```

## slicing array
```go
numbers := []int{0,1, 2, 3}
one = numbers[0:1]   //[1]

// a slice expression is formed by specifying a start or a low bound and a stop or high bound like  a[start:stop].
    // this selects a range of elements which includes the element at index start, but excludes the element at index stop.
 
    // slicing an array returns a slice, not an array

```
 