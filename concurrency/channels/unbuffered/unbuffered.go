package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// channels block if there is not receiver waiting on it.
	// so launcing a go routine for sending value to channel.
	// After lauching go routine, main routine goes and waits on
	// receiver end. This makes sure sender and receiver sending and
	// receiving synchronously

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
