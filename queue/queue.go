package main

import (
	"fmt"
)

type Queue struct {
	front int
	elements []int
}

func (q *Queue)Enqueue(item int) {
	q.elements = append(q.elements, item)

}

func (q *Queue)Dequeue() int{
	if q.IsEmptyQueue() {
		return 0
	}

	ele := q.elements[q.front]
	q.front++
	q.elements = q.elements[q.front:len(q.elements)]

	return ele
}

func (q *Queue)IsEmptyQueue() bool{
	return q.front == len(q.elements)-1
}
func CreateQueue() *Queue {
	q := &Queue{}
	q.elements = make([]int, 0)
	q.front = 0
	return q
}

func main() {
	q := CreateQueue()
	q.Enqueue(10)
	q.Enqueue(11)
	q.Enqueue(12)
	q.Enqueue(13)
	
	q.Enqueue(14)
	fmt.Println(*q, len(q.elements))

	q.Dequeue()
	fmt.Println(*q, len(q.elements))

	q.Dequeue()
	fmt.Println(*q, len(q.elements))

	q.Enqueue(10)
	fmt.Println(*q, len(q.elements))

}
