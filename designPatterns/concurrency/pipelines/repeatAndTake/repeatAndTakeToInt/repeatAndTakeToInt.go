package main

import (
	"fmt"
	"math/rand"
	"time"
)

type isPrime struct {
	num     int
	status  bool
	divisor int
}

func main() {
	prime := func(num int) isPrime {
		i := 2
		v := isPrime{num, true, num}
		for ; i < num/2; i++ {
			if num%i == 0 {
				v = isPrime{num, false, i}
				break
			}
		}
		return v
	}

	primeFinder := func(done <-chan interface{}, valueStream <-chan interface{}, Num int) <-chan interface{} {
		primeStream := make(chan interface{})
		go func() {
			defer close(primeStream)
			for i := 0; i < Num; i++ {
				v := <-valueStream
				select {
				case <-done:
					return
				case primeStream <- prime(v.(int)):
				}
			}

		}()

		return primeStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case takeStream <- v:
				}
			}

		}()
		return takeStream
	}

	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}

			}
		}()
		return valueStream
	}

	done := make(chan interface{})
	defer close(done)

	timeTaken := time.Now()

	rand := func() interface{} { return rand.Intn(10000) }
	for v := range take(done, primeFinder(done, repeatFn(done, rand), 1000)) {
		fmt.Println(v)
	}

	fmt.Println("Total time taken to execute: ", time.Since(timeTaken))
}
