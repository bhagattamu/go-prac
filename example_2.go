// Parameter passing

package main

import "fmt"

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) int {
	return n * 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)
	val := 10
	double(val)
	fmt.Println(val) // 10 pass by value
	doublePtr(&val)
	fmt.Println(val) // 20 pass by reference using pointer
}
