package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type homePageSize struct {
	url  string
	size int
}

func main() {
	urls := []string{
		"https://www.amazon.in",
		"https://www.apple.com",
		"https://www.google.com",
		"https://www.microsoft.com",
	}

	results := make(chan homePageSize)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("Error")
				panic(err)
			}

			results <- homePageSize{
				url:  url,
				size: len(bs),
			}
		}(url)
	}

	fmt.Println("routines are scheduled... ")
	var biggest homePageSize

	for i, _ := range urls {
		result := <-results
		fmt.Println(i, "results:", result)
		if result.size > biggest.size {
			biggest = result
		}
	}

	fmt.Println("Biggest gome page is: ", biggest.url)
}
