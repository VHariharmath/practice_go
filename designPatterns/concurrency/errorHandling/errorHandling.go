package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Err  error
	Resp *http.Response
}

func main() {
	checkStatus := func(done chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				resp, err := http.Get(url)
				result := Result{Err: err, Resp: resp}

				select {
				case <-done:
					return
				case results <- result:
				}
			}

		}()
		return results
	}

	var errCount int
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com/", "a", "b", "c", "d"}

	for result := range checkStatus(done, urls...) {
		if result.Err != nil {
			fmt.Println("Error: ", result.Err)
			errCount++
			if errCount > 3 {
				fmt.Println("Too many errors, breaking")
				break
			}
			continue
		}
		fmt.Printf("Response: %v \n", result.Resp.Status)
	}
}
