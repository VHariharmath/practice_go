package main

import (
	"fmt"
)

func deleteOneElementFromSlice(slice ...int) {
	oldSlice := slice
	slice = oldSlice[0 : len(oldSlice)-1]
}
func incrementByOneToElements(slice ...int) {

	for i := 0; i < len(slice); i++ {
		slice[i]++
	}
}
func main() {
	var buffer [256]int
	slice := buffer[20:30]

	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}

	fmt.Println("Orginal slice:\t\t\t\t", slice, "\t\tlen:\t", len(slice))

	// adding one to each element in a slice
	incrementByOneToElements(slice...)
	fmt.Println("After increment:\t\t\t", slice, "\tlen:\t", len(slice))

	// deleting an element from a slice
	deleteOneElementFromSlice(slice...)
	fmt.Println("After deleting an element:\t\t", slice, "\tlen:\t", len(slice))

	// cut and join 2 slices
	newSlice := []int{11, 12, 13, 14, 15}
	slice = append(slice[0:2], newSlice...)
	fmt.Println("New slice:\t\t\t\t", slice, "\t\tlen:\t", len(slice))

	//Slicing a slice
	slice = slice[1:5]
	fmt.Println("After slicing a slice:\t\t\t", slice, "\t\t\tlen:\t", len(slice))

}
