package main

import (
	"fmt"
	"math"
)

type rect struct {
	l float64
	b float64
}

type circ struct {
	r float64
}

type measure interface {
	area() float64
}

func geometry(m measure) float64 {
	return m.area()
}

func (a rect) area() float64 {
	return a.l * a.b
}

func (b circ) area() float64 {
	return math.Pi * b.r * b.r
}

func main() {
	x := rect{10, 20}
	y := circ{10}
	fmt.Println(geometry(x))
	fmt.Println(geometry(y))
}
