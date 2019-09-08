package main

import "fmt"

var y = 42
var z = "Shaken, not stirred"
var a = `James said, "Shaken, not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	// z = 43 // Cannot assign int value to variable of type string
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)

	// Types have Zero values or default values
	/*
	- false for booleans
	- 0 for integers
	- 0.0 for floats
	- "" for strings
	- nil for
		- pointers
		- functions
		- interfaces
		- slices
		- channels
		- maps
	- use short declaration operator as much as possible
	- use var for
		- zero value
		- package scope

	*/
}
