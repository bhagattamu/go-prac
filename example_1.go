// Slice is sequence of items

package main

import (
	"fmt"
)

func main() {
	// Same type
	loons := []string{"bugs", "daffy", "tazz"}

	fmt.Printf("%v type of %T \n", loons, loons)

	// length
	fmt.Println(len(loons)) // 3

	// 0 indexing
	fmt.Println(loons[1]) // daffy

	// slices
	fmt.Println(loons[1:]) // [daffy tazz]
}
