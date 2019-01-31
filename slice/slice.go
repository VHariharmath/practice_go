package main

import (
	"fmt"
)

func deleteOneElementFromSlice(slice *[]byte) {
	sliceptr := *slice
	*slice = sliceptr[0 : len(sliceptr)-1]
}
func incrementByOneToElements(slice []byte) {

	for i := 0; i < len(slice); i++ {
		slice[i]++
	}
}
func main() {
	var buffer [256]byte
	slice := buffer[20:30]

	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	fmt.Println(slice)

	incrementByOneToElements(slice)

	fmt.Println(slice)

	fmt.Println("Before:\t", slice, len(slice))
	deleteOneElementFromSlice(&slice)
	fmt.Println("After:\t", slice, len(slice))

}
