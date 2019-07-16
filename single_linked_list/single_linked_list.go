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

func (listhead *listHead) findMiddleNode() *node {
	if listhead.head == nil {
		return nil
	}

	slowPtr := listhead.head
	fastPtr := listhead.head

	for fastPtr != nil && fastPtr.next != nil {
		fastPtr = fastPtr.next.next
		slowPtr = slowPtr.next
	}

	return slowPtr
}

func (listhead *listHead) detectAndRemoveLoop() bool {
	if listhead.head == nil {
		fmt.Println("No loop")
		return false
	}
	slowPtr := listhead.head
	fastPtr := listhead.head
	for slowPtr != nil && fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			fmt.Println("Loop node = ", slowPtr.data)
			removeLoop(slowPtr, listhead.head)
			return true
		}

	}

	return false
}

func removeLoop(loopNode *node, listhead *node) {

	ptr1 := listhead
	var ptr2 *node
	for true {
		ptr2 = loopNode

		for ptr2.next != loopNode && ptr2.next != ptr1 {
			ptr2 = ptr2.next
		}

		if ptr2.next == ptr1 {
			fmt.Println(ptr2.data)
			break
		}

		ptr1 = ptr1.next
	}

	ptr2.next = nil

}

func (listhead *listHead) deleteNode(key int) {
	root := listhead.head

	if root == nil {
		fmt.Println("List is empty")
		return
	}

	if root.data == key {
		listhead.head = root.next
		return
	}

	var prev *node

	for root != nil && root.data != key {
		prev = root
		root = root.next
	}

	prev.next = root.next
}

func (listhead *listHead) deleteRear() {
	root := listhead.head
	if root == nil {
		fmt.Println("List is empty")
		return
	}

	var prev *node
	for root.next != nil {
		prev = root
		root = root.next
	}

	prev.next = nil

}

func (listhead *listHead) deleteFront() {
	root := listhead.head
	if root == nil {
		fmt.Println("List is empty")
		return
	} else {
		listhead.head = root.next
	}
}

func main() {
	ll := &listHead{}
	ll.insertTail(6)
	ll.insertTail(7)
	ll.insertTail(8)
	ll.insertTail(9)
	ll.insertFront(5)

	fmt.Println("Elements in the list")
	ll.printList()

	/*
		fmt.Println("Middile element = ", ll.findMiddleNode().data)

		ll.head.next.next.next.next.next = ll.head.next.next

		fmt.Println(ll.detectAndRemoveLoop())
		ll.printList()

		ll.deleteNode(8)

		ll.deleteRear()
	*/
	ll.deleteFront()
	fmt.Println("after deleting ")
	ll.printList()
}
