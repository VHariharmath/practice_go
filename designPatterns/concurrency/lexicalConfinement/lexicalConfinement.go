package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	printer := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buf bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buf, "%c", b)
		}
		fmt.Println(buf.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	// constraining the goroutines to access only subset is lexical confinement here.
	// printer() cannot access other elements which are not passed. So the lexical
	// scope not allows any wrong thing. So need of any synchronization mechanism here
	go printer(&wg, data[:3])
	go printer(&wg, data[3:])
	wg.Wait()

}
