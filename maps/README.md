# This sections descirbe the usage and implementaion of maps in golang

```go

    //declaration
    var people map[string]string

    // initialisation 
     
     employees := map[string]string{}

     people = make(map[string]string)

     balances := map[string]float64{
        "USD":111.5,
        "EUR":10000,
     }

     v,ok := balances["CHF"]
     if ok {
        fmt.Println("CHF balance exsists",v)
     }else{
        fmt.Println("Given balance doesnot exsist")
     }
```

## modifying the values or adding more values into map
```go
 
 balances := map[string]float64{
        "USD":111.5,
        "EUR":10000,
        "GBP":15,
}

balances["INR"] = 10000

for k, v := range balances {
		fmt.Println(k, v)
	}

delete(balances, "GBP")

```


## Comparing maps
```go

// Comapring String maps
    a := map[string]string{"A": "0", "B": "1"}
	b := map[string]string{"Y": "24", "Z": "25"}
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are  not equal")    // => Not equal 
	}


```

## modifications to maps

Maps when copied they have same reference
```go
   people := map[string]string{"Naruto":"male","Sasuke":"male","Hinata":"female"}
   clone := people 
   
   clone["Sakura"]="female"

/// !!!!!!!!!! the people map is also modified

   fmt.Println(people) // => map[Hinata:female Naruto:male Sakura:female Sasuke:male]

// CLONING map
   clone = make(map[string]string)
   for k,v := range people{
    clone[k]=v
   }
   clone["Jiraya"]="male"
   fmt.Println(people,clone) 
```


### How to clone a map?
1. initialize a new map
2. use a for loop to copy each element into the new map