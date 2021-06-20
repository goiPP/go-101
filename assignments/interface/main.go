package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}
type square struct {
	side float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {

	sq := square{5}
	tri := triangle{2, 5}

	printArea(sq)
	printArea(tri)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (sq square) getArea() float64 {
	return math.Pow(sq.side, 2)
}

func (tri triangle) getArea() float64 {
	return 0.5 * tri.height * tri.base
}
