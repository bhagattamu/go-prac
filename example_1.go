// Basic function definition

package main

import "fmt"

// Return sumation of a and b
func add(a int, b int) int {
	return a + b
}

// divmod returns quotient and remainder
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	val := add(1, 2)
	fmt.Println(val)
	div, mod := divmod(1, 2)
	fmt.Printf("div = %d mod = %d \n", div, mod)
}
