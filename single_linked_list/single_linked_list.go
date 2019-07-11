package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type listHead struct {
	head *node
}

func (listhead *listHead) insertTail(key int) {

	item := &node{key, nil}

	if listhead.head == nil {
		listhead.head = item
		return
	}

	var prev *node
	for i := listhead.head; i != nil; i = i.next {
		prev = i
	}
	prev.next = item
}

func (listhead *listHead) printList() {
	for i := listhead.head; i != nil; i = i.next {
		fmt.Println(i.data)
	}
}

func (listhead *listHead) insertFront(key int) {
	item := &node{key, nil}

	if listhead.head == nil {
		listhead.head = item
		return
	}
	item.next = listhead.head
	listhead.head = item
}

func main() {
	ll := &listHead{}
	ll.insertTail(6)
	ll.insertTail(7)
	ll.insertTail(8)
	ll.insertTail(9)
	ll.insertFront(5)
	ll.insertFront(4)
	ll.insertFront(3)
	ll.insertFront(2)
	ll.insertFront(1)

	ll.printList()
}
