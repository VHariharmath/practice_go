package main

import (
	"fmt"
)

func main() {
	e := make(chan int)
	o := make(chan int)
	q := make(chan bool)
	// send routine
	go foo(e, o, q)

	//receive routine, not necessarily a go routine
	bar(e, o, q)
}

func foo(e, o chan int, q chan bool) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	// close will send signal over the q channel
	close(q)
}

func bar(e, o <-chan int, q <-chan bool) {

	for {
		select {
		case x := <-e:
			fmt.Println("Response to eve channel: ", x)
		case y := <-o:
			fmt.Println("Response from odd channel: ", y)
		case z, ok := <-q:
			if !ok {
				fmt.Println("Response from quit channel, break the loop : ", z, ok)
				return
			} else {
				fmt.Println("Response from quit channel ", ok)
			}
			return
		}
	}
}
