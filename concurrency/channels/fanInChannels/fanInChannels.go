package main

import (
	"fmt"
	"sync"
)

func main() {
	e := make(chan int)
	o := make(chan int)
	// channel to fanning in the result
	fanIn := make(chan int)
	// send routine
	go foo(e, o)

	//receive routine, not necessarily a go routine
	go bar(e, o, fanIn)

	for i := range fanIn {
		fmt.Println(i)
	}
}

func foo(e, o chan int) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func bar(e, o <-chan int, fanIn chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := range e {
			fanIn <- i
		}
		wg.Done()
	}()

	go func() {
		for i := range o {
			fanIn <- i
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn)
}
