package main

import (
	"fmt"
	"runtime"
)

func main() {
	generator := func(done <-chan interface{}, values ...int) <-chan int {
		intstream := make(chan int)
		go func() {
			defer close(intstream)
			for _, i := range values {
				select {
				case <-done:
					return
				case intstream <- i:
				}
			}
		}()
		return intstream
	}
	add := func(done <-chan interface{}, values <-chan int, additive int) <-chan int {
		addstream := make(chan int)
		go func() {
			defer close(addstream)
			for v := range values {
				select {
				case <-done:
					return
				case addstream <- (v + additive):
				}
			}

		}()
		return addstream
	}

	multiply := func(done <-chan interface{}, values <-chan int, multiplier int) <-chan int {
		multiplyStream := make(chan int)
		go func() {
			defer close(multiplyStream)
			for v := range values {
				select {
				case <-done:
					return
				case multiplyStream <- (v * multiplier):
				}
			}

		}()
		return multiplyStream
	}

	done := make(chan interface{})
	defer close(done)
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	intStream := generator(done, s...)
	pipeline := multiply(done, add(done, multiply(done, intStream, 1), 2), 3)

	for i := range pipeline {
		fmt.Println(i)
	}

	fmt.Println("Num of Goroutines: ", runtime.NumGoroutine())
}
