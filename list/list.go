package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	x.PushBack(12)
	x.PushFront(11)
	x.PushBack(13)
	x.PushFront(10)

	var el *list.Element
	count := 0

	for i := x.Front(); i != nil; i = i.Next() {
		count++
		if count == 3 {
			el = i
		}
		fmt.Println(i.Value)
	}

	fmt.Println()
	x.InsertAfter(20, el)

	for i := x.Front(); i != nil; i = i.Next() {

		fmt.Println(i.Value)
	}
}
