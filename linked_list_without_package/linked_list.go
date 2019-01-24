package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func (head *Node) insertTail(key int) *Node {

	item := new(Node)
	item.data = key
	var prev *Node

	if head == nil {
		return item
	} else {
		for i := head; i != nil; i = i.next {
			prev = i
		}
		prev.next = item
		return head
	}
}

func printList(head *Node) {
	for i := head; i != nil; i = i.next {
		fmt.Println(i.data)
	}
}

func (head *Node) insertFront(key int) *Node {
	item := new(Node)
	item.data = key

	if head == nil {
		return item
	} else {
		item.next = head
		return item
	}
}

func (head *Node) insertAfter(index, key int) *Node {
	var count int
	item := new(Node)
	item.data = key

	cur := head
	for ; cur != nil; cur = cur.next {
		count++
		if index == count {
			break
		}
	}

	var temp *Node

	temp = cur.next
	cur.next = item
	item.next = temp

	return head
}

func (head *Node) insertBefore(index, key int) *Node {

	var count int
	item := new(Node)
	item.data = key

	cur := head
	var prev *Node
	for ; cur != nil; cur = cur.next {
		count++
		if index == count {
			break
		}
		prev = cur
	}

	var temp *Node

	temp = prev.next
	prev.next = item
	item.next = temp

	return head
}

func main() {
	head := new(Node)
	head = head.insertTail(6)
	head = head.insertTail(7)
	head = head.insertTail(8)
	head = head.insertTail(9)
	head = head.insertFront(5)
	head = head.insertFront(4)
	head = head.insertFront(3)
	head = head.insertFront(2)
	head = head.insertFront(1)

	head = head.insertAfter(3, 100)
	head = head.insertBefore(3, 200)
	printList(head)
}
