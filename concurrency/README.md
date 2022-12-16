## Go Routines
- A goroutine is a lightweight thread of execution; goroutines are key ingredients to achieve concurrency in Go.
- A goroutine is a function that is capable of running concurrently with other functions. To create a goroutine we use the keyword go followed by a function invocation;
- Goroutines are far smaller that threads, they typically take around 2kB of stack space to initialize compared to a thread which takes a fixed size of 1-2Mb.
- An OS Thread Stack is fixed size but a goroutine stack size shrinks and grows as needed.
- Scheduling a goroutine is much cheaper than scheduling a thread.
- OS threads are scheduled by the OS kernel, but goroutines are scheduled by its own Go Scheduler using a technique called m:n scheduling, because it multiplexes (or schedules) m goroutines on n OS threads.
- Goroutines have no identity. There is no notion of identity that is accessible to the programmer.

## Go routine and wait groups

```go
func printStr(str string){
    fmt.Println("Hello ",str)
}
func main(){
     names:=[]string{"Naruto","Sasuke","Sakura"}
     for _,name := range names{
        go printStr(name)   // ==> "go"  is used to make it a go routine
     }
     time.Sleep(time.Second * 1) // ==> sleep is used because if the main exits the go routine will not be executed anymore
}
```

## modification to above program to use waitgroups

```go
func printStr(str string, wg *sync.WaitGroup){
    fmt.Println("Hello ",str)
    wg.Done()   // ==> Done to indicate the wait group is done
}
func main(){
   
     names:=[]string{"Naruto","Sasuke","Sakura"}
     var wg sync.WaitGroup
     wg.Add(len(names))   // No.of wait groups needed in the program
     for _,name := range names{
        go printStr(name,&wg)   // ==> "go"  is used to make it a go routine
     }
     wg.Wait() // ==> Wait  to indicate the Main to wait till the go routines are completed 
}
```

## To detect data race(mutiple go routines modifying a value at same time) use "-race" arguemnt while running the go program
```
go run -race main.go
```

## To avoid concurrency in go we can use mutexes to lock and unlock the variables to prevent simultaneous write
```go
n:=100
var m sync.Mutex
func Increment(wg  *sync.WaitGroup){
    m.Lock()  // ==> locked 
    n++
    m.UnLock()  // ==> unlocked
    wg.Done()
}
func Decrement(wg *sync.WaitGroup){
    m.Lock()    // ==> locked 
    n--
    m.UnLock()  // ==> unlocked
    wg.Done()
}
func main(){
    var wg sync.WaitGroup
    wg.add(200)
    for i:=0;i<=100;i++ {
        Increment(&wg)
        Decrement(&wg)
    }
    wg.Wait()
}

```

## Go channels
A Go channel is a communication mechanism that allows Goroutines to exchange data. When developers have numerous Goroutines running at the same time, channels are the most convenient way to communicate with each other.

Developers often use these channels for notifications and managing concurrency in applications.

`
channels can also be uni-directional
only_receving := make(<-chan string)
only_sending := make(chan <- string)
`
use defer statement for closing channels , best practices
`
use defer close(channel) to postpone the closing of the channel
`