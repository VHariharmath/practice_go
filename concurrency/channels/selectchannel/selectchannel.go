package main

import (
	"fmt"
)

func main() {
	e := make(chan int)
	o := make(chan int)
	q := make(chan int)
	// send routine
	go foo(e, o, q)

	//receive routine, not necessarily a go routine
	bar(e, o, q)
}

func foo(e, o, q chan<- int) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}

func bar(e, o, q <-chan int) {

	for {
		select {
		case x := <-e:
			fmt.Println("Response to eve channel: ", x)
		case y := <-o:
			fmt.Println("Response from odd channel: ", y)
		case z := <-q:
			fmt.Println("Response from quit channel: ", z)
			return
		}
	}
}
