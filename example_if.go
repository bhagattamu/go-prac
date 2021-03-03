// Example of if statement

package main

import (
	"fmt"
)

func main() {
	// assigning var x with 10, it will give type integer
	x := 10
	// single if
	if x > 5 {
		fmt.Println("Big number")
	}
	// if and else
	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}
	// and
	if x > 5 && x < 15 {
		fmt.Println("x is right")
	}
	// or
	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}
	a := 11.0
	b := 20.0
	// if can have optional initialization
	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}
