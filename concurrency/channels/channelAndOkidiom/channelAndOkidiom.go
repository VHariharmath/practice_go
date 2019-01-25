package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 1947
		close(c)
	}()

	v, ok := <-c
	fmt.Println("channel value: ", v, "channe open ? ", ok)

	v, ok = <-c
	fmt.Println("channel value: ", v, "channel open ? ", ok)
}
