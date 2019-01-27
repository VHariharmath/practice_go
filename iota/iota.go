package main

import (
	"fmt"
)

const (
	p = iota
	q
	r
	s
)

const (
	x1 = iota + 1
	x2
	x3
	x4
)

const (
	y1 = iota + 1
	y2 = iota + 2
	y3 = iota + 3
	y4 = iota + 4
)
const (
	_ = iota
	z1
	z2
	z3
	z4
)

const (
	kb = 1 << ((iota + 1) * 10)
	mb = 1 << ((iota + 1) * 10)
	gb = 1 << ((iota + 1) * 10)
	tb = 1 << ((iota + 1) * 10)
)

const (
	a = (iota + 1) * 1024
	b
	c
	d
)

const (
	year2020 = iota + 2020
	year2021
	year2022
	year2023
)

func main() {

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)

	fmt.Println()

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)

	fmt.Println()

	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)

	fmt.Println()

	fmt.Println(z1)
	fmt.Println(z2)
	fmt.Println(z3)
	fmt.Println(z4)

	fmt.Println()

	fmt.Printf("%b\n", kb)
	fmt.Printf("%b\n", mb)
	fmt.Printf("%b\n", gb)
	fmt.Printf("%b\n", tb)

	fmt.Println()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println()

	fmt.Println(year2020)
	fmt.Println(year2021)
	fmt.Println(year2022)
	fmt.Println(year2023)
}
