package main

import (
	"fmt"
)

type Queue struct {
	elements []int
}

type Stack struct {
	elements []int
}

type AbstractDataType interface {
	list(item int)
	unlist() int
}

func (q *Queue)list(item int) {
	q.elements = append(q.elements, item)

}

func (q *Queue)unlist() int{
	if q.IsEmptyQueue() {
		return 0
	}

	ele := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]

	return ele
}

func (q *Queue)IsEmptyQueue() bool{
	return len(q.elements) == 0
}
func CreateQueue() *Queue {
	q := &Queue{}
	q.elements = make([]int, 0)
	return q
}

func (s *Stack)list(item int) {
	s.elements = append(s.elements, item)
}

func (s *Stack)unlist() int {

	if s.IsEmpty() {
		return 0
	}

	len := len(s.elements)
	ele := s.elements[len-1]
	s.elements = s.elements[:len-1]

	return ele
}

func (s *Stack)IsEmpty() bool {
	return len(s.elements) == 0 
}

func CreateStack() *Stack {
	s := &Stack{}
	s.elements = make([]int, 0)
	return s
}

func listElement(adt AbstractDataType, item int) {
	adt.list(item)
}

func unlistElement(adt AbstractDataType) int {
	return adt.unlist()
}


func main() {
	q := CreateQueue()

	fmt.Println("Queue operatations using interface")
	listElement(q, 10)
	listElement(q, 11)
	listElement(q, 12)
	listElement(q, 13)
	
	listElement(q, 14)
	fmt.Println(*q, len(q.elements))

	unlistElement(q)
	fmt.Println(*q, len(q.elements))

	unlistElement(q)
	fmt.Println(*q, len(q.elements))

	unlistElement(q)
	fmt.Println(*q, len(q.elements))

	listElement(q, 15)
	fmt.Println(*q, len(q.elements))

	s := CreateStack()
	fmt.Println("Stack operations using interface")

	listElement(s, 1)
	listElement(s, 2)
	listElement(s, 3)
	listElement(s, 4)

	fmt.Println(*s, len(s.elements))	
	
	unlistElement(s)
	fmt.Println(*s, len(s.elements))

	unlistElement(s)
	fmt.Println(*s, len(s.elements))

	listElement(s, 5)
	fmt.Println(*s, len(s.elements))
}