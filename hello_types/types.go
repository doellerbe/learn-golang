package main

import "fmt"

type hotdog int

// Example creating a type
func main() {
	b := hotdog(12)
	fmt.Println(b)

	// Convert a type
	a := int(b)
	fmt.Println(a)

	// assign a new value to a
	a = 9999
	fmt.Println(a)

	// a = hotdog(20) // cannot assign to type hotdog anymore!
}
