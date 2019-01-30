package main

import (
	"fmt"
)

func main() {
	take := func(done <-chan interface{}, valueStream <-chan string, Num int) <-chan string {
		takeStream := make(chan string)
		go func() {
			defer close(takeStream)
			for i := 0; i < Num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}

		}()
		return takeStream
	}

	repeatFn := func(done <-chan interface{}, values ...string) <-chan string {
		valueStream := make(chan string)
		go func() {
			defer close(valueStream)
			for {
				for _, str := range values {
					select {
					case <-done:
						return
					case valueStream <- str:
					}
				}

			}

		}()
		return valueStream
	}

	done := make(chan interface{})
	defer close(done)
	var message string
	for v := range take(done, repeatFn(done, "Hello", "How", "are", "you", "?"), 5) {
		message += v
	}
	fmt.Println(message)
}
