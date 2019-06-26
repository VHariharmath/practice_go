package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack)Push(item int) {
	s.elements = append(s.elements, item)
}

func (s *Stack)Pop() int {

	if s.IsEmpty() {
		return 0
	}

	len := len(s.elements)
	ele := s.elements[len-1]
	s.elements = s.elements[:len-1]

	return ele
}

func (s *Stack)IsEmpty() bool {
	return len(s.elements) == 0 
}

func main() {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Printf("Stack size : %v\n", len(s.elements))
	fmt.Println("Stack elements")
	fmt.Println(*s)

	fmt.Println("Popped element")
	fmt.Println(s.Pop())

	fmt.Printf("Stack size : %v\n", len(s.elements))
	fmt.Println("Stack elements")
	fmt.Println(*s)

	fmt.Println("Popped element")
	fmt.Println(s.Pop())

	fmt.Printf("Stack size : %v\n", len(s.elements))
	fmt.Println("Stack elements")
	fmt.Println(*s)

	s.Push(10)
	fmt.Printf("Stack size : %v\n", len(s.elements))
	fmt.Println("Stack elements")
	fmt.Println(*s)

}