package main

import "fmt"

func arrayRotateLeft(a []int, pos int) []int {

	for i:= 0 ; i < pos; i++ {
		rotateLeft(a)
	}

	return a
}

func rotateLeft(a []int) {
	var i int
	temp := a[0]

	for i=0;i<len(a)-1; i++ {
		a[i] = a[i+1]
	}

	a[i] = temp
}
func main() {
	a := []int{1,2,3,4,5,6,7,8}
	fmt.Println(a)

	fmt.Println("Array after rotating 4 elements to the left")
	fmt.Println(arrayRotateLeft(a, 4))
}