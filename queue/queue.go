package main

import (
	"fmt"
)

type Queue struct {
	elements []int
}

func (q *Queue)Enqueue(item int) {
	q.elements = append(q.elements, item)

}

func (q *Queue)Dequeue() int{
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
