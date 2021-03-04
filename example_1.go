// Go strings

package main

import (
	"fmt"
)

func main() {
	book := "The color of magic"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Printf("%v, type of %T \n", book[0], book[0])
	// uint8 is a byte
	// Strings in go are immutable
	// book[0] = 112 -> throws error

	// slice (start, end), 0 based, 1/2 empty range
	fmt.Println(book[4:11]) // expected: color o
	// slice no end
	fmt.Println(book[4:]) // expected: color of magic
	// slice no start
	fmt.Println(book[:4]) // expected: The

	// Use + to concatenate strings

	fmt.Println("t" + book[1:])

	// Unicode

	fmt.Println("It was ½ price.")

	// Multi line

	poem := `
	The road goes ever on
	Down from the door where it began
	...
	`
	fmt.Println(poem)
}
