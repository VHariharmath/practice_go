package main

import (
	"fmt"
	"sync"
)

type Button struct {
	clicked *sync.Cond
}

func main() {

	button := Button{clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goRoutineRunning sync.WaitGroup
		goRoutineRunning.Add(1)
		go func() {
			goRoutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goRoutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.clicked, func() {
		fmt.Println("Maximizing window")
		clickRegistered.Done()
	})

	subscribe(button.clicked, func() {
		fmt.Println("Displaying i/p box")
		clickRegistered.Done()
	})
	subscribe(button.clicked, func() {
		fmt.Println("Mouse clicked")
		clickRegistered.Done()
	})

	button.clicked.Broadcast()
	clickRegistered.Wait()
}
