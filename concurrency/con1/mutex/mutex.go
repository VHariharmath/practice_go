package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	const gs = 100
	var mu sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			counter++
			runtime.Gosched()
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
