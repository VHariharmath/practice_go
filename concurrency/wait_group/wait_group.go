package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo() is called: ", i, "times")
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar() is called: ", i, "times")
	}
}

func main() {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())

	go foo()
	wg.Add(1)
	bar()

	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	wg.Wait()
}
