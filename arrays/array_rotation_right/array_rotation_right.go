package main

import "fmt"

func arrayRightRotation(a []int, pos int) []int {

	for i:= 0; i < pos; i++ {
		rightRotate(a)
	}

	return a
}

func rightRotate(a []int) {
	temp := a[len(a)-1]
	
	i := len(a)-1

	for ; i > 0; i-- {
		a[i] = a[i-1]
	}

	a[i] = temp
	
}
func main() {
	a := []int{1,2,3,4,5,6,7,8}
	fmt.Println(a)
	fmt.Println("Array after rotating 4 postion to right")
	fmt.Println(arrayRightRotation(a, 4))
}