package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// declaration is important as we are calling inside anonymous function.
	// declaration will make compiler know that or method exists
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):

				}
			}
		}()
		return orDone
	}

	sig := func(duration time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(duration)
		}()
		return c
	}

	start := time.Now()
	fmt.Println("Num go routines before:", runtime.NumGoroutine())
	<-or(sig(1*time.Hour), sig(2*time.Minute), sig(1*time.Minute), sig(1*time.Hour), sig(1*time.Second))
	//This exits after a second. When second timer writes something on the channel, the last stack of channels
	// will be closed because of deferred close on orDone channel. Our main routine is blocked on this channel.
	// After the channel close(), blocked main routine continues and exits. However there will be go routines leak
	// which needs to be handled.
	fmt.Println("Exited after:", time.Since(start), "seconds time")
	fmt.Println("Num go routines:", runtime.NumGoroutine())
}
