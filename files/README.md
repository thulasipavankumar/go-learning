## Creating and geting info of the file

```go
   // os.Create() function creates a file if it doesn't already exist. If it exists, the file is truncated.

    newFile, err := os.Create("new.txt")

    // TRUNCATING A FILE
    err = os.Truncate("new.txt", 0) //0 means completely empty the file.

    // CLOSING THE FILE
    newFile.Close()
    
    // OPEN AND CLOSE AN EXISTING FILE
    file, err := os.Open("a.txt") // open in read-only mode

     //OPENING a FILE WITH MORE OPTIONS
    file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)
    // We can Use opening attributes individually or combined
    // using an OR between them
    // e.g. os.O_CREATE|os.O_APPEND
    // or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
    // os.O_RDONLY // Read only
    // os.O_WRONLY // Write only
    // os.O_RDWR // Read and write
    // os.O_APPEND // Append to end of file
    // os.O_CREATE // Create is none exist
    // os.O_TRUNC // Truncate file when opening

 var fileInfo os.FileInfo
    fileInfo, err = os.Stat("a.txt")
 
    p := fmt.Println
    p("File Name:", fileInfo.Name())        // => File Name: a.txt
    p("Size in bytes:", fileInfo.Size())    // => Size in bytes: 0
    p("Last modified:", fileInfo.ModTime()) // => Last modified: 2019-10-21 16:16:00.325037748 +0300 EEST
    p("Is Directory? ", fileInfo.IsDir())   // => Is Directory?  false
    p("Pemissions:", fileInfo.Mode())       // => Pemissions: -rw-r-----
 
    // CHECKING IF FILE EXISTS
    fileInfo, err = os.Stat("b.txt")
    // error handling
    if err != nil {
        if os.IsNotExist(err) {
            log.Fatal("The file does not exist")
        }
    }
 
    // RENAMING AND MOVING A FILE
    oldPath := "a.txt"
    newPath := "aaa.txt"
    err = os.Rename(oldPath, newPath)
    // error handling
    if err != nil {
        log.Fatal(err)
    }
 
    // REMOVING A FILE
    err = os.Remove("aa.txt")
    // error handling
    if err != nil {
        log.Fatal(err)
    }

```

## Reading using scanio package and scanner with delimiters 
```go
// opening the file in read-only mode. The file must exist (in the current working directory)
    // use a valid path!
    file, err := os.Open("my_file.txt")
    // error handling
    if err != nil {
        log.Fatal(err)
    }
    // defer closing the file
    defer file.Close()
 
    // the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
    scanner := bufio.NewScanner(file)
 
    // the default scanner is bufio.ScanLines and that means it will scan a file line by line.
    // there are also bufio.ScanWords and bufio.ScanRunes.
    // scanner.Split(bufio.ScanLines)
 
    // scanning for next token in this case \n which is the line delimiter.
    success := scanner.Scan() //read a line
    if success == false {
        // false on error or EOF. Check for errors
        err = scanner.Err()
        if err == nil {
            log.Println("Scan was completed and it reached End Of File.")
        } else {
            log.Fatal(err)
        }
    }
 
    // Getting the data from the scanner with Bytes() or Text()
    fmt.Println("First Line found:", scanner.Text())
    //If we want the next token, so the next line or \n, we call scanner.Scan() again
 
    // Reading the whole remaining part of the file:
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
 
    // Checking for any possible errors:
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
```

## Reading from STDIN
```go
    p := log.Println
	pf := log.Printf
	scanner := bufio.NewScanner(os.Stdin)
	p(scanner.Scan())
	pf("%s\n", scanner.Text())
	pf("%v\n", scanner.Bytes())
```