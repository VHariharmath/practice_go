package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Get the context of main Go routine
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Before starting context:", runtime.NumGoroutine(), ctx.Err())
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("parent go routine cancelled")
				return
			default:
				n++
				time.Sleep(200 * time.Millisecond)
				fmt.Println("Working: ", n)
			}
		}
	}()

	// This wait will give time for go routine to run for sometime
	time.Sleep(2 * time.Second)
	fmt.Println("Before cancel", runtime.NumGoroutine(), ctx.Err())
	cancel()
	// wait for sometime to see the context value reflected
	time.Sleep(2 * time.Second)
	fmt.Println("After cancel", runtime.NumGoroutine(), ctx.Err())
}
