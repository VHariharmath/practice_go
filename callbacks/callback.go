package main

import "fmt"

func sum(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

func even(f func(args ...int) int, args ...int) int {
	xi := []int{}
	for _, v := range args {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}

	return f(xi...)
}

func odd(f func(args ...int) int, args ...int) int {
	xi := []int{}
	for _, v := range args {
		if v%2 == 1 {
			xi = append(xi, v)
		}
	}

	return f(xi...)
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := sum(x...)
	fmt.Println(s1)
	s2 := even(sum, x...)
	fmt.Println(s2)
	s3 := odd(sum, x...)
	fmt.Println(s3)

}
