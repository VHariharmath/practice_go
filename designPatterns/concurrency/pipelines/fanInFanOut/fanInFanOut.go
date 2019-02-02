package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type isPrime struct {
	num     int
	status  bool
	divisor int
}

func main() {
	prime := func(num int) isPrime {
		v := isPrime{num, true, num}
		for i := 2; i < num/2; i++ {
			if num%i == 0 {
				v = isPrime{num, false, i}
				break
			}
		}
		return v
	}

	primeFinder := func(done <-chan interface{}, valueStream <-chan int, Num int) <-chan isPrime {
		primeStream := make(chan isPrime)
		go func() {
			defer close(primeStream)
			for i := 0; i < Num; i++ {
				v := <-valueStream
				select {
				case <-done:
					return
				case primeStream <- prime(v):
				}
			}
		}()
		return (primeStream)
	}

	fanOut := func(done <-chan interface{}, valueStream <-chan int, num int) []<-chan isPrime {
		numFinders := runtime.NumCPU()
		fmt.Println("Number of CPUs:", num)
		primeStream := make([]<-chan isPrime, numFinders)
		for i := 0; i < numFinders; i++ {
			primeStream[i] = primeFinder(done, valueStream, num)
		}
		return primeStream
	}

	fanIn := func(done <-chan interface{}, channels []<-chan isPrime) <-chan isPrime {
		var wg sync.WaitGroup
		multiplexedStream := make(chan isPrime)

		multiplex := func(done <-chan interface{}, c <-chan isPrime) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(done, c)
		}

		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	repeatFn := func(done <-chan interface{}, fn func() int) <-chan int {
		valueStream := make(chan int)
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

	start := time.Now()
	rand := func() int { return rand.Intn(1000000) }
	randStream := repeatFn(done, rand)
	primeStream := fanIn(done, fanOut(done, randStream, 10000))
	for v := range primeStream {
		fmt.Printf("%v is %v and divisor = %v\n", v.num, v.status, v.divisor)
	}
	fmt.Println("Time taken: ", time.Since(start))
}
