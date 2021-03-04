package main

import (
	"fmt"
)

// Continue to move next loop instead of running below code
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("before continue", i)
		if i > 5 {
			continue
		}
		fmt.Println("after continue", i)
	}
}
