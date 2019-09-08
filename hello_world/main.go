package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	foo()

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	// declare and assign a variable with :=
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)
}

func foo() {
	fmt.Println("In foo")
}
