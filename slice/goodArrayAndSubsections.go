package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 0, 1, 0, 1, 2, 3, 4}
	k := 5

	fmt.Println("Original array")
	fmt.Println(a)
	sections := len(a)/ k
	newArraylength := len(a) - len(a) % k

	fmt.Println("New good array")
	a = a[0:newArraylength]
	fmt.Println(a)

	subSections := make([][]int, sections)

	for i:=0; i<sections; i++ {
		lowerIndex := newArraylength/sections * i
		higherIndex := newArraylength/sections * (i+1)
		subSections[i] = a[lowerIndex:higherIndex]
	}

	
	fmt.Printf("New good array length = %v and number of subsections can be %v \n", len(a), sections)
	fmt.Println("Subsections of new good array")
	for i:=0; i < sections; i++ {
		fmt.Println(subSections[i])
	}
}