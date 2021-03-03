// Calculate mean

package main

import (
	"fmt"
)

func main() {
	x := 2.0
	y := 2.9

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)
	// Or can write mean := (x + y) / 2
	var mean float64
	mean = (x + y) / 2
	fmt.Printf("result=%v, type of %T\n", mean, mean)
}
