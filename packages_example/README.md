## Two types of packages
1. Executable packages (the name is pre-defined it's main) 
2. Non-Executable packages

## File scope , local scope and global scope
```go
import f "fmt"
func main(){
    f.Println("Hello World!")
}
```
## init function execute before main - there can be mutiple init function in a go file