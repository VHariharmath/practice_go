package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	const gs = 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
}
