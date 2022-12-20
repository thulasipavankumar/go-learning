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

## Go module dependencies are saved in $GOPATH/pkg/mod

## Commands - Create and Publish a Module on GitHub
### Creating Your Own Go Module


Consider the following steps to create your own Go module and publish it on GitHub.

Create a free GitHub account if you don’t have one.

1. Using the browser, create a new repository on GitHub

e.g. go_math



2. Create a directory for the module anywhere on the disk. Inside it create a directory for each package of the module.

The name of the directory will be the name of the module on GitHub.

Write the code for each package.

Don’t forget to export all names! That means that the first letter of each name must be uppercase.



3. The next step is to initialize the module by running go mod init and the module path from the module directory. This will generate go.mod file that stores the import path and any dependencies.

e.g. go mod init github.com/ddadumitrescu/go_math



The executed commands:

$mkdir go_math

$cd go_math/

$mkdir calc geometry

$go mod init github.com/ddadumitrescu/go_math

go: creating new go.mod: module github.com/ddadumitrescu/go_math

$cat go.mod

module github.com/ddadumitrescu/go_math



go 1.13

$cd ..

$ls

go_math

$tree .

.

└── go_math

├── calc

├── geometry

└── go.mod



3 directories, 1 file



Publish the Module on GitHub


Move to the module directory (e.g. cd go_math)

1. Initialize the local module folder as a git repository.

git init

2. Add the remote repository and give it the name origin

git remote add origin https://github.com/ddadumitrescu/go_math.git

Check the name and the URL of the remote repository: git remote -v

3. Add all files from the current directory to the staging area.

git add -A

4. Set some variables

git config user.name "andrei"

git config user.email "someone@someplace.com"

5. Commit the changes

git commit -m "some init msg"

6. Synchronize the local and remote repositories.

git push -u -f origin master

Authenticating …

Now the local and the remote repositories are synchronized.

7. Version the module or make the first release.

Create a git tag:

git tag -a v1.0.0 -m "initial release"

git push origin master --tags

