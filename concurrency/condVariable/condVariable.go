package main

import (
	"fmt"
	"sync"
	"time"
)

type emptyStruct struct {
}

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	// create slice of length = 0 and capacity = 10
	queue := make([]interface{}, 0, 10)

	removeFromTheQueue := func(delay time.Duration) {
		time.Sleep(delay)
		cond.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from the queue")
		cond.L.Unlock()
		cond.Signal()
	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		if len(queue) == 2 {
			cond.Wait()
		}
		fmt.Println("Added to the queue")
		queue = append(queue, emptyStruct{})
		go removeFromTheQueue(1 * time.Second)
		cond.L.Unlock()
	}
}
