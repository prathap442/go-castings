package main

import "fmt"

//we just can initialize a variable outside of the main function but we cannot use it in the main function in the go language

func main() {
	// We might be able to initialize a variable and then later assign a variable to it.  Is the following code valid?
	deckSize := 20
	fmt.Println(deckSize)
}