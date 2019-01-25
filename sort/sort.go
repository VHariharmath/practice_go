package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	x := []person{{"Virat", 30}, {"AB", 31}, {"Ajinkya", 29}}

	sort.Sort(ByAge(x))

	fmt.Println(x)

}
