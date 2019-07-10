package main

import "fmt"

func binarySearch(a []int, key int, start int, end int) (bool, int) {
	if end < start {
		return false, -1
	}

	mid := start + (end-start)/2

	if a[mid] == key {
		return true, mid
	}

	if a[mid] > key {
		return binarySearch(a, key, start, mid-1)
	} else {
		return binarySearch(a, key, mid+1, end)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 10, 11, 12, 13, 14, 15, 21, 22, 23}

	found, pos := binarySearch(a, 5, 0, len(a)-1)
	fmt.Println("found ?", found, "at position = ", pos)

	found, pos = binarySearch(a, 111, 0, len(a)-1)
	fmt.Println("found ?", found, "at position = ", pos)

	found, pos = binarySearch(a, 23, 0, len(a)-1)
	fmt.Println("found ?", found, "at position = ", pos)

	found, pos = binarySearch(a, 1, 0, len(a)-1)
	fmt.Println("found ?", found, "at position = ", pos)
}
