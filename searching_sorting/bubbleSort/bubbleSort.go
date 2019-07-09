package main

import "fmt"

func bubbleSort(a []int) []int {

	len := len(a)

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}

func main() {
	a := []int{10, 20, 5, 15, 2, 3, 4, 11, 12, 13}

	fmt.Println(a)

	fmt.Println(bubbleSort(a))

}
