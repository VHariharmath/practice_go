package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	kb = 1 << ((iota + 1) * 10)
	mb = 1 << ((iota + 1) * 10)
	gb = 1 << ((iota + 1) * 10)
	tb = 1 << ((iota + 1) * 10)
)

func main() {
	var wg sync.WaitGroup
	c := make(<-chan int)
	ctx, cancel := context.WithCancel(context.Background())

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	noop := func() {
		wg.Done()
		select {
		case <-ctx.Done():
			return
		}
		<-c
	}
	NumGoroutines := 100000
	wg.Add(NumGoroutines)
	fmt.Println("Memory Consumed before: ", memConsumed()/mb, "mb")
	for i := NumGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	fmt.Println("Memory consumed after: ", memConsumed()/mb, "mb", "Num of go routines", runtime.NumGoroutine())

	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("Memory consumed after cancel(): ", memConsumed()/mb, "mb", "Num of go routines", runtime.NumGoroutine())
}
