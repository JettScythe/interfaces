package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{5, 7}
	sq := square{5}
	printArea(t)
	printArea(sq)
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
