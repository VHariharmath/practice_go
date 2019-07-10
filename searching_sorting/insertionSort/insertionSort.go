package main

import "fmt"

func insertionSort(a []int) []int {
	len := len(a)

	for i := 1; i < len; i++ {
		key := a[i]
		j := i - 1

		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j = j - 1
		}

		a[j+1] = key
	}

	return a
}

func main() {
	a := []int{10, 20, 5, 15, 2, 3, 4, 11, 12, 13}

	fmt.Println(a)

	fmt.Println(insertionSort(a))

}
