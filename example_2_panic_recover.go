package main

import (
	"fmt"
)

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}()
	return vals[index]
}

func main() {
	/*
		comment out section with pani
		c := []int{1, 2, 3}
		fmt.Println(c[10])
	*/
	/*
		file, err := os.Open("no-such-file")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Println("file opened")
	*/
	v := safeValue([]int{1, 2, 3}, 10)
	fmt.Println(v)
}
