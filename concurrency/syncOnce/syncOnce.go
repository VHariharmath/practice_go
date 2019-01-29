package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	increment := func() {
		count++
	}

	var increments sync.WaitGroup
	var once sync.Once

	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			// Do will be called only once irrespective of how many times Do is called
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Println("count:", count)
}
