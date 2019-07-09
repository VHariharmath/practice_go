package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func selectionSort(a []int) []int {
	len := len(a)

	for i := 0; i < len-1; i++ {
		minIdx := i
		for j := i; j < len-1; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}

		swap(&a[minIdx], &a[i])
	}

	return a
}

func main() {
	a := []int{10, 20, 5, 15, 2, 3, 4, 11, 12, 13}

	fmt.Println(a)

	fmt.Println(selectionSort(a))

}
