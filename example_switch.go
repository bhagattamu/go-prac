// example of switch statement

package main

import (
	"fmt"
)

func main() {
	x := 2

	switch x {
	case 1:
		fmt.Println("One")
		break
	case 2:
		fmt.Println("Two")
		break
	case 3:
		fmt.Println("Three")
		break
	default:
		fmt.Println("Default")
	}

	// Switch without expression we have condition in case

	switch {
	case x > 100:
		fmt.Println("x greater than 100")
		break
	case x > 10:
		fmt.Println("x greater than 10 but less than 100")
		break
	case x > 5:
		fmt.Println("x greater than 5 but less than 10")
		break
	default:
		fmt.Println("x less than 5")
	}
}
