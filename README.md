# The Go Lang:
  - The Compiler based language developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson 
  - The Go lang is a compiler based language like C but with the better ways of the garbage collection, memory usage and nack way of optmising the resources entity .
  - File extensions do end with the `.go`

# First program let us go with the Hello World
  - As indicated in this section we have file1.go which is the hello world program
```
    package main

    //package main is being used whenever we do try to build some executables
    import "fmt"

    func main() {
        fmt.Println("Hello World!")
    }
```
  - Here there is a code written in go lang and let us see the types of the codes in go 

  ```
    Two types of Packages in go
    - Executable: Writing the Executable go code(using the Package Main)
    - Reusable: Writing the Reusable go code(using  the custom package)
  ```
  Writing the Executable code tries to make use of the package main .
  Writing the reusable code doesnot do that .



# The func main is being used here 