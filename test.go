package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Square struct {
	center Point
	length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	point := Point{x, y}
	if length <= 0 {
		return nil, fmt.Errorf("Length must be >= 0 (was %d)", length)
	}
	square := &Square{
		center: point,
		length: length,
	}
	return square, nil
}

func (s *Square) move(dx int, dy int) {
	s.center.x += dx
	s.center.y += dy
}

func (s *Square) Area() int {
	return s.length * s.length
}

func main() {
	square, err := NewSquare(8, 8, 10)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("%+v\n", square)
	fmt.Printf("Area: %d\n", square.Area())
	fmt.Println("Moving by 2,2")
	square.move(2, 2)
	fmt.Printf("%+v\n", square)
}
