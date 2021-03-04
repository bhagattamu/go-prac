/*
Print the number from 1 to 20.
However if the number is divisible by 3 then it should print "fizz" instead of number,
if the number is divisible by 5 then it should print "buzz",
if the number is divisible by both then it should print "fizz buzz"
*/

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
