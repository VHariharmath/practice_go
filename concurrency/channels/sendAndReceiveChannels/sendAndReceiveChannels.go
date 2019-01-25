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
	c <- 1947
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
