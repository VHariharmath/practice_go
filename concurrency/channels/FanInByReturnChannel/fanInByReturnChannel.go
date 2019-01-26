package main

import (
	"fmt"
)

func main() {
	// channel to fanning in the result
	c := fanIn(foo("Virat"), foo("Kohli"))

	for i := 0; i < 2; i++ {
		fmt.Println(i, <-c)
	}
}

// takes receive channels as parameter and returns receive channel
func fanIn(s1, s2 <-chan string) <-chan string {
	fanIn := make(chan string)
	go func() {
		fanIn <- <-s1
	}()

	go func() {
		fanIn <- <-s2
	}()

	return fanIn
}

// returns receive channel
func foo(str string) <-chan string {
	c := make(chan string)
	go func() {
		c <- str
	}()

	return c
}
