package main

import (
	"fmt"
	"time"
)

func main() {
	// channel to fanning in the result
	c := fanIn(foo("Virat"), foo("Kohli"))

	// receiving for n times, here n is 10, as send and receive is cont loops, n is our wish here
	for i := 0; i < 10; i++ {
		fmt.Println(i, <-c)
	}
}

// takes receive channels as parameter and returns receive channel
func fanIn(s1, s2 <-chan string) <-chan string {
	fanIn := make(chan string)
	go func() {
		for {
			fanIn <- <-s1
		}
	}()

	go func() {
		for {
			fanIn <- <-s2
		}
	}()

	return fanIn
}

// returns receive channel
func foo(str string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%v %v\n", i, str)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	return c
}
