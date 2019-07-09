package main

import "fmt"

func linearSearch(a []int, key int) (bool, int) {
	for i, v := range a {
		if v == key {
			return true, i
		}
	}

	return false, -1
}

func main() {
	a := []int{10, 20, 5, 15, 2, 3, 4, 11, 12, 13}

	found, pos := linearSearch(a, 3)
	fmt.Println("found ?", found, "at position = ", pos)

	found, pos = linearSearch(a, 100)
	fmt.Println("found ?", found, "at position = ", pos)
}
