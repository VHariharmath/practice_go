package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// send routine
	go foo(c)

	//receive routine, not necessarily a go routine
	bar(c)
}

func foo(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}

	//closing is important, otherwise 'range' wont see the values
	// comment below close() call, run this and see what happens
	close(c)
}

func bar(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
