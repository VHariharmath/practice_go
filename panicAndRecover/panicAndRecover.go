package main

import "fmt"

func main() {
	foo()
	fmt.Println("Returned from foo()")
}

func foo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in foo()", r)
		}
	}()
	fmt.Println("Calling bar")
	bar(0)
	fmt.Println("Returned from bar()")
}

func bar(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Deffering in bar(): ", i)
	fmt.Println("Printing in bar(): ", i)
	bar(i + 1)
}
