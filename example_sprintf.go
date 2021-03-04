package main

import (
	"fmt"
)

func main() {
	i := 42
	fmt.Printf("Printf => i=%v, type of %T \n", i, i)
	s := fmt.Sprintf("%d", i)
	fmt.Printf("Sprintf => s=%s, type of %T \n", s, s)
	// %q add quotes to the string
	fmt.Printf("Sprintf => s=%q, type of %T \n", s, s)
}
