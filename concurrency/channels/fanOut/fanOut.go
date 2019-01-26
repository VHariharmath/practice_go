package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populateOders(c1)

	go fanOut(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
}

func populateOders(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func fanOut(in <-chan int, out chan<- int) {
	var wg sync.WaitGroup

	// Run through the range and fan out the work
	for i := range in {
		wg.Add(1)
		go func(n int) {
			out <- generateOrdernum(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(out)
}

func generateOrdernum(n int) int {
	return n + rand.Intn(1000)
}
