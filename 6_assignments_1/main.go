package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 2.5, height: 5}
	s := square{sideLength: 13}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base * t.height)
}
func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}

func printArea(s shape) {
	fmt.Println("Shape area:", s.getArea())
}
