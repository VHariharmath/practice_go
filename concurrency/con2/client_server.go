package main

import (
	"fmt"
)

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(args []int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}

	return sum
}

func server(req chan Request) {

	fmt.Println("Inside server: ")
	request := <-req

	fmt.Println(request)
	for _, v := range request.args {
		fmt.Println(v)
	}

	sum := request.f(request.args)
	fmt.Println("Sending response to client")
	request.resultChan <- sum
}

func main() {
	ChannelRequests := make(chan Request)
	request := Request{[]int{10, 20, 30}, sum, make(chan int)}

	go func() {
		fmt.Println("Sending client request: ")
		ChannelRequests <- request
	}()
	go server(ChannelRequests)

	fmt.Println("Response from server ", <-request.resultChan)

}
