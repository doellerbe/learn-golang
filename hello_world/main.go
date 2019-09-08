package main

import "fmt"

var z int // assigned value of 0
var w = 999

func main() {
	fmt.Println("Hello, World!")

	foo()

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	// declare and assign a variable with :=
	x := 99
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)

	// use var declared outside of function body
	z = 42
	fmt.Println(z)
	fmt.Println(w)
}

func foo() {
	fmt.Println("In foo")
}
