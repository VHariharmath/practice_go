package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type Linkedlist struct {
	headNode *Node
}

func (ll *Linkedlist) nodeBetweenValues(item1 int, item2 int) *Node {
	var cur *Node
	for cur = ll.headNode; cur != nil; cur = cur.next {
		if cur.prev != nil && cur.next != nil {
			if (cur.prev.data == item1 && cur.next.data == item2) ||
				(cur.prev.data == item2 && cur.next.data == item1) {
				return cur
			}
		}
	}

	return nil
}

func (ll *Linkedlist) insertFront(item int) {
	newNode := &Node{}
	newNode.data = item

	if ll.headNode != nil {
		newNode.next = ll.headNode
		ll.headNode.prev = newNode
	}

	ll.headNode = newNode
}

func (ll *Linkedlist) insertTail(item int) {
	newNode := Node{}
	newNode.data = item

	if ll.headNode == nil {
		ll.headNode = &newNode
		return
	}

	v := ll.headNode
	for ; v.next != nil; v = v.next {
	}

	v.next = &newNode
	newNode.prev = v
}

func (ll *Linkedlist) display() {
	for v := ll.headNode; v != nil; v = v.next {
		fmt.Println(v.data)
	}
}

func (ll *Linkedlist) reverseList() {

	var temp, cur *Node
	for cur = ll.headNode; cur != nil; {
		temp = cur.prev
		cur.prev = cur.next
		cur.next = temp
		cur = cur.prev
	}

	ll.headNode = temp.prev
}

func (ll *Linkedlist) lastNode() *Node {
	var cur *Node

	for cur = ll.headNode; cur.next != nil; cur = cur.next {

	}

	return cur
}
func main() {
	var ll Linkedlist
	ll.insertTail(101)
	ll.insertTail(102)
	ll.insertTail(103)
	ll.insertTail(104)
	ll.insertFront(100)
	fmt.Println("List after adding elements")
	ll.display()

	ll.reverseList()
	fmt.Println("After reversing the ll")
	ll.display()

	fmt.Println("Last node")
	fmt.Println(ll.lastNode().data)

	fmt.Println("Node between 103, 101")
	fmt.Println(ll.nodeBetweenValues(103, 101))

	fmt.Println("Node between 101, 103")
	fmt.Println(ll.nodeBetweenValues(101, 103))

	fmt.Println("Node between 103 and 105")
	fmt.Println(ll.nodeBetweenValues(103, 105))
}
