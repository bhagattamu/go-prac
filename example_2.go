package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "tazz"}

	// for
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	// Single value range
	for i := range loons {
		fmt.Println(i)
	}

	// Double value range
	for i, name := range loons {
		fmt.Printf("%s %d \n", name, i)
	}

	// Double value range ignore index by using _
	// note: Unused variable in a go are a compilation error
	for _, name := range loons {
		fmt.Println(name)
	}

	// extend slice / append
	loons = append(loons, "elmer")
	fmt.Println(loons) // [bugs daffy tazz elmer]

}
