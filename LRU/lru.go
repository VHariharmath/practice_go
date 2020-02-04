package main

import "fmt"

type node struct {
	pageNumber int
	prev       *node
	next       *node
}
type que struct {
	numberOfPages int
	count         int
	front         *node
	rear          *node
}

type hash struct {
	cap   int
	array []*node
}

func createQ(cap int) *que {
	q := new(que)
	if q == nil {
		panic("No memory")
	}

	q.front = nil
	q.rear = nil
	q.numberOfPages = cap
	return q
}

func createHash(cap int) *hash {
	h := new(hash)
	if h == nil {
		panic("No memory")
	}

	h.cap = cap
	h.array = make([]*node, cap)

	return h
}

func areAllFramesFull(q *que) bool {
	return q.count == q.numberOfPages
}

func noFramesFound(q *que) bool {
	return q.rear == nil
}

func enque(q *que, h *hash, pn int) {
	if areAllFramesFull(q) == true {
		h.array[q.rear.pageNumber] = nil
		deque(q)
	}

	newPage := new(node)
	if newPage == nil {
		panic("No memory found")
	}
	newPage.next = q.front

	if noFramesFound(q) == true {
		q.front = newPage
		q.rear = newPage
	} else {
		q.front.prev = newPage
		q.front = newPage
	}
	h.array[pn] = newPage
	q.count++
}

func deque(q *que) {
	if noFramesFound(q) == true {
		return
	}

	if q.front == q.rear {
		q.front = nil
	}

	q.rear = q.rear.prev
	if q.rear != nil {
		q.rear.next = nil
	}

	q.count--

}

func referencePage(q *que, h *hash, pn int) {
	reqPage := h.array[pn]
	if reqPage == nil {
		enque(q, h, pn)
	} else if reqPage != q.front {
		reqPage.prev.next = reqPage.next
		if reqPage.next != nil {
			reqPage.next.prev = reqPage.prev
		}

		if reqPage == q.rear {
			q.rear = reqPage.prev
			q.rear.next = nil
		}

		reqPage.next = q.front
		reqPage.prev = nil
		q.front.prev = reqPage
		q.front = reqPage
	}

}

func main() {
	q := createQ(4)
	h := createHash(10)

	referencePage(q, h, 1)
	referencePage(q, h, 2)
	referencePage(q, h, 3)
	referencePage(q, h, 1)
	referencePage(q, h, 4)
	referencePage(q, h, 5)

	fmt.Println(q.front.pageNumber)
	fmt.Println(q.front.next.pageNumber)
	fmt.Println(q.front.next.next.pageNumber)
	fmt.Println(q.front.next.next.next.pageNumber)

}
