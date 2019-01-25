package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	// buffered channels wont block until it finds the resp channel
	// is filled. Its asynchronous in nature. Send and move on.

	c <- 42

	fmt.Println(<-c)
}
