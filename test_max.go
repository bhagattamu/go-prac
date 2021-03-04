package main

import "fmt"

func main() {
	nums := []int{4, 6, 8, 3, 68, 9}
	max := nums[0]
	for _, value := range nums[1:] {
		if max < value {
			max = value
		}
	}
	fmt.Println(max)
}
